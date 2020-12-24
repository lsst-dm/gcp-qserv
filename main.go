package main

// BEFORE RUNNING:
// ---------------
// 1. If not already done, enable the Kubernetes Engine API
//    and check the quota for your project at
//    https://console.developers.google.com/apis/api/container
// 2. This sample uses Application Default Credentials for authentication.
//    If not already done, install the gcloud CLI from
//    https://cloud.google.com/sdk/ and run
//    `gcloud beta auth application-default login`.
//    For more information, see
//    https://developers.google.com/identity/protocols/application-default-credentials
// 3. Install and update the Go dependencies by running `go get -u` in the
//    project directory.

import (
	"flag"
	"log"
)

func main() {

	project := flag.String("project", "qserv-dev-3d7e", "Identifier of GCP project")

	region := "us-central1"
	cluster := "qserv-dev"
	czarNodePool := "czar-pool-b1d5"
	workerNodePool := "worker-pool-ab14"

	sizeWorkerPtr := flag.Int64("num-workers", 5, "Number of workers")
	sizeCzarPtr := flag.Int64("num-czars", 1, "Number of czars")
	sizeDown := flag.Bool("downsize", false, "Delete all cluster nodes, take precendence over other size options")

	flag.Parse()

	// create GCloud Client
	gcloud, err := NewGCloudClient(cluster, *project, region)
	if err != nil {
		log.Fatal("Error creating GCloud client")
	}

	gcloudContainer, err := gcloud.NewGCloudContainerClient()
	if err != nil {
		log.Fatal("Error creating GCloudContainer client")
	}

	sizeCzar := *sizeCzarPtr
	sizeWorker := *sizeWorkerPtr
	if *sizeDown {
		sizeCzar = 0
		sizeWorker = 0
	}

	log.Printf("Resizing czar node pool to %d", sizeCzar)
	gcloudContainer.SetNodePoolSize(czarNodePool, sizeCzar)
	log.Printf("Resizing worker node pool to %d", sizeWorker)
	gcloudContainer.SetNodePoolSize(workerNodePool, sizeWorker)
}
