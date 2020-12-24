package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/google"
	container "google.golang.org/api/container/v1"
)

const (
	// operationWaitTimeoutSecond define the time wait in second before assuming the failure of a GCloud operation
	operationWaitTimeoutSecond = 600

	// operationPollIntervalSecond define the interval in second before each GCloud operation status check
	operationPollIntervalSecond = 10
)

type GCloud struct {
	Client  *http.Client
	Cluster string
	Context context.Context
	Project string
	Region  string
}

type GCloudClient interface {
	NewGCloudContainerClient() (GCloudContainerClient, error)
}

// NewGCloudClient return a GCloud client
func NewGCloudClient(cluster string, project string, region string) (gcloud GCloudClient, err error) {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, container.CloudPlatformScope)

	if err != nil {
		err = fmt.Errorf("Error creating GCloud client:\n%v", err)
	}

	gcloud = &GCloud{
		Client:  client,
		Cluster: cluster,
		Context: ctx,
		Project: project,
		Region:  region,
	}

	return
}

// NewGCloudContainerClient return a GCloud container client
func (g *GCloud) NewGCloudContainerClient() (gcloud GCloudContainerClient, err error) {
	service, err := container.New(g.Client)

	if err != nil {
		err = fmt.Errorf("Error creating GCloud container client:\n%v", err)
		return
	}

	gcloud = &GCloudContainer{
		Client:  g,
		Service: service,
	}

	return
}
