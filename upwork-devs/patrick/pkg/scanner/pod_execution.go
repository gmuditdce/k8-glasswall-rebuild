package scanner

import (
	"context"
	"os"
	"path/filepath"
	"strconv"

	kcorev1 "k8s.io/api/core/v1"
	kmetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// This method will run the scan again'st a file encapsulated the job.
// It starts a pod, passing it the path to the folder where the file is
// TODO : Copy that file to a location where the pod can access it. we will use minio for a first test
func (job Job) processFile(kubeClient client.Client, processSettings *ProcessSettings) {

	log.Info("Processing file : " + job.Filename)

	task_id_str := strconv.Itoa(job.TaskID)

	// This is the location where will be the input file will be copied to, for processing
	processing_location := filepath.Join(processSettings.ProcessingFolder, job.Batch+"-"+task_id_str)
	output_location := filepath.Join(processSettings.OutputFolder, job.Batch+"-"+task_id_str)

	// Need to create the processing folder where the file will be moved to
	os.MkdirAll(processing_location, os.ModePerm)
	os.MkdirAll(output_location, os.ModePerm)

	from := filepath.Join(processSettings.SourceFolder, job.Filename)
	to := filepath.Join(processing_location, job.Filename)

	//We copy the source file to his processing folder
	err := moveFileToProcessingFolder(from, to)
	if err != nil {
		log.Error(err, "Could not copy the file, worker exiting without processing")
		// TODO Notification about the failure
		return
	}

	pod := &kcorev1.Pod{
		ObjectMeta: kmetav1.ObjectMeta{
			Name:      "rebuild-" + task_id_str + "-" + job.Batch,
			Namespace: job.Namespace,
		},
		Spec: kcorev1.PodSpec{
			Volumes: []kcorev1.Volume{
				{
					Name: "file-to-process",
					VolumeSource: kcorev1.VolumeSource{
						PersistentVolumeClaim: &kcorev1.PersistentVolumeClaimVolumeSource{
							ClaimName: "rebuild-pvc",
						},
					},
				},
			},
			InitContainers: []kcorev1.Container{
				{
					Name:  "rebuild",
					Image: "azopat/gw-rebuild",
					Env: []kcorev1.EnvVar{
						{Name: "INPUT_LOCATION", Value: processing_location},
						{Name: "OUTPUT_LOCATION", Value: output_location},
					},
					VolumeMounts: []kcorev1.VolumeMount{{Name: "file-to-process", MountPath: "/data"}},
				},
			},
			Containers: []kcorev1.Container{
				{
					Name:  "upload-to-storage",
					Image: "azopat/minio",
					Env: []kcorev1.EnvVar{
						{Name: "STORAGE_ENDPOINT", Value: processSettings.StorageEndpoint},
						{Name: "STORAGE_ACCESS_KEY", Value: processSettings.StorageAccessKey},
						{Name: "STORAGE_SECRET_KEY", Value: processSettings.StorageSecretKey},
						{Name: "OUTPUT_FOLDER", Value: output_location},
						{Name: "STORAGE_BUCKET", Value: processSettings.StorageBucket},
					},
					VolumeMounts: []kcorev1.VolumeMount{{Name: "file-to-process", MountPath: "/data"}},
				},
			},
			RestartPolicy: kcorev1.RestartPolicyNever,
		},
	}

	err = kubeClient.Create(context.Background(), pod)
	if err != nil {
		log.Error(err, "Could not create the pod")
	}

	// TODO: check pod execution result and notify about the failure

}

func moveFileToProcessingFolder(from string, to string) error {

	err := os.Rename(from, to)

	if err != nil {
		return err
	}

	return nil

}
