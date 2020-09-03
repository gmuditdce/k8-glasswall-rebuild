package scanner

import (
	"io/ioutil"

	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("scan_processor")

type ScanProcessor struct {
	Batch          string
	Folder         string
	ContainerImage string // Image containing scan tool - https://github.com/filetrust/Glasswall-Rebuild-SDK-Evaluation
	Namespace      string // the namespace where the pods will be created
}

func (s *ScanProcessor) ScanFiles() {

	log.Info("Scan processor execution", "info", s)
	files, err := ioutil.ReadDir(s.Folder)
	if err != nil {
		log.Error(err, "Could not read the directory")
	}

	i := 1
	for _, f := range files {
		if !f.IsDir() {
			job := Job{Filename: f.Name(), TaskID: i, Batch: s.Batch, ContainerImage: s.ContainerImage, Namespace: s.Namespace}
			log.Info("Adding a new job to the queue", "info", job)
			i = i + 1
			JobQueue <- job
		}
	}

}
