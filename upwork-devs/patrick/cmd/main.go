package main

import (
	"os"

	"github.com/google/uuid"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"

	"github.com/filetrust/Open-Source/upwork/project-k8-glasswall-rebuild/pkg/scanner"
	"github.com/urfave/cli/v2"
)

const (
	FLAG_SOURCE_FOLDER     = "source_folder"
	FLAG_PROCESSING_FOLDER = "processing_folder"
	FLAG_OUTPUT_FOLDER     = "output_folder"

	FLAG_IMAGE     = "rebuild_image"
	FLAG_NAMESPACE = "namespace"

	FLAG_STORAGE_ACCESS_KEY = "storage_access_key"
	FLAG_STORAGE_SECRET_KEY = "storage_secret_key"
	FLAG_STORAGE_BUCKET     = "storage_bucket"
	FLAG_STORAGE_ENDPOINT   = "storage_endpoint"

	FLAG_MAX_QUEUE  = "max_queue"
	FLAG_MAX_WORKER = "max_worker"
)

var log = logf.Log.WithName("cmd")

func main() {

	// The logger instantiated here can be changed to any logger
	// implementing the logr.Logger interface. This logger will
	// be propagated through the whole operator, generating
	// uniform and structured logs.
	logf.SetLogger(zap.Logger(false))

	cliApp := &cli.App{}
	cliApp.Action = runIt
	cliApp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    FLAG_SOURCE_FOLDER,
			Usage:   "Source folder where controller watch files",
			Value:   "/data/src-files",
			EnvVars: []string{"SOURCE_FOLDER"},
		},
		&cli.StringFlag{
			Name:    FLAG_PROCESSING_FOLDER,
			Usage:   "Files being processed",
			Value:   "/data/processing-files",
			EnvVars: []string{"PROCESSING_FOLDER"},
		},
		&cli.StringFlag{
			Name:    FLAG_OUTPUT_FOLDER,
			Usage:   "Folder where processed files are moved",
			EnvVars: []string{"OUTPUT_FOLDER"},
		},
		&cli.StringFlag{
			Name:    FLAG_IMAGE,
			Usage:   "Glasswall rebuild engine container image",
			Value:   "azopat/gw-rebuild",
			EnvVars: []string{"IMAGE"},
		},
		&cli.StringFlag{
			Name:    FLAG_NAMESPACE,
			Usage:   "Namespace where to spawn worker pods",
			Value:   "test",
			EnvVars: []string{"NAMESPACE"},
		},
		&cli.StringFlag{
			Name:    FLAG_STORAGE_ACCESS_KEY,
			Usage:   "Destination storage access key",
			Value:   "minio",
			EnvVars: []string{"STORAGE_ACCESS_KEY"},
		},
		&cli.StringFlag{
			Name:    FLAG_STORAGE_SECRET_KEY,
			Usage:   "Destination storage secret key",
			Value:   "minio123",
			EnvVars: []string{"STORAGE_SECRET_KEY"},
		},
		&cli.StringFlag{
			Name:    FLAG_STORAGE_BUCKET,
			Usage:   "Destination storage bucket name",
			Value:   "glasswall",
			EnvVars: []string{"STORAGE_BUCKET"},
		},
		&cli.StringFlag{
			Name:    FLAG_STORAGE_ENDPOINT,
			Usage:   "Destination storage endpoint",
			Value:   "http://minio.default.svc.cluster.local:9000",
			EnvVars: []string{"STORAGE_ENDPOINT"},
		},
		&cli.IntFlag{
			Name:    FLAG_MAX_QUEUE,
			Usage:   "Queue channel capacity",
			Value:   5,
			EnvVars: []string{"MAX_QUEUE"},
		},
		&cli.IntFlag{
			Name:    FLAG_MAX_WORKER,
			Usage:   "Maximum number of workers processing files",
			Value:   5,
			EnvVars: []string{"MAX_WORKER"},
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Error(err, "Could not start execution")
	}

}

func runIt(c *cli.Context) error {

	// Get a kubernetes client
	k8sclient, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		log.Error(err, "Kubernetes config not found, exiting")
		os.Exit(1)
	}
	log.Info("Connected to the cluster")

	// Watch and processing folders
	sourceFolder := c.String(FLAG_SOURCE_FOLDER)
	processingFolder := c.String(FLAG_PROCESSING_FOLDER)
	outputFolder := c.String(FLAG_OUTPUT_FOLDER)

	// COntainer images
	image := c.String(FLAG_IMAGE)
	namespace := c.String(FLAG_NAMESPACE)

	// Storage details
	storageAccessKey := c.String(FLAG_STORAGE_ACCESS_KEY)
	storageSecretKey := c.String(FLAG_STORAGE_SECRET_KEY)
	storageBucket := c.String(FLAG_STORAGE_BUCKET)
	storageEndpoint := c.String(FLAG_STORAGE_ENDPOINT)

	// This should be injected as env variables, represent amount of simultaneaus job
	maxQueue := c.Int(FLAG_MAX_QUEUE)
	maxWorker := c.Int(FLAG_MAX_WORKER)

	processSettings := &scanner.ProcessSettings{SourceFolder: sourceFolder, ProcessingFolder: processingFolder, ProcessPodImage: image, ProcessPodNamespace: namespace, OutputFolder: outputFolder, StorageAccessKey: storageAccessKey, StorageSecretKey: storageSecretKey, StorageBucket: storageBucket, StorageEndpoint: storageEndpoint}
	log.Info("Starting controller with settings ", "info", processSettings)

	// Workers initialization
	scanner.JobQueue = make(chan scanner.Job, maxQueue)
	dispatcher := scanner.NewDispatcher(maxWorker, k8sclient, processSettings)
	dispatcher.Run()
	log.Info("Job dispatcher started")

	// Starting a scan
	scanProcessor := scanner.ScanProcessor{Folder: sourceFolder, Batch: uuid.New().String(), ContainerImage: image, Namespace: namespace}
	scanProcessor.ScanFiles()

	// This is just to keep the pod running for now.
	for {

	}

}
