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
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/container/v1"
)

func main() {

	projectPtr := flag.String("project", "qserv-dev-3d7e", "Identifier of GCP project")

	region := "us-central1"
	cluster := "qserv-dev"
	czarNodePool := "czar-pool-b1d5"
	workerNodePool := "worker-pool-ab14"

	sizeWorkerPtr := flag.Int64("num-workers", 5, "Number of workers")
	sizeCzarPtr := flag.Int64("num-czars", 1, "Number of czars")
	sizeDown := flag.Bool("downsize", false, "Delete all cluster nodes, take precendence over other size options")

	flag.Parse()

	ctx := context.Background()

	c, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	containerService, err := container.New(c)
	if err != nil {
		log.Fatal(err)
	}

	// The name (project, location, cluster, node pool id) of the node pool to set
	// size.
	// Specified in the format 'projects/*/locations/*/clusters/*/nodePools/*'.

	clusterName := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", *projectPtr, region, cluster)

	sizeCzar := *sizeCzarPtr
	sizeWorker := *sizeWorkerPtr
	if *sizeDown {
		sizeCzar = 0
		sizeWorker = 0
	}

	fmt.Printf("Resizing czar node pool to %d", sizeCzar)
	resizeNodePool(ctx, clusterName, czarNodePool, containerService, sizeCzar)
	fmt.Printf("Resizing worker node pool to %d", sizeWorker)
	resizeNodePool(ctx, clusterName, workerNodePool, containerService, sizeWorker)

}

func resizeNodePool(ctx context.Context, clusterName string, nodePoolName string,
	containerService *container.Service, size int64) {

	name := fmt.Sprintf("%s/nodePools/%s", clusterName, nodePoolName)
	rb := &container.SetNodePoolSizeRequest{
		// TODO: Add desired fields of the request body.
	}
	rb.NodeCount = size

	resp, err := containerService.Projects.Locations.Clusters.NodePools.SetSize(name, rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("%#v\n", resp)
}
