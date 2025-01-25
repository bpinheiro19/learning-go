package main

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

var (
	resourceGroupClient *armresources.ResourceGroupsClient
)

func createResourceGroupClient(subscriptionId string, credential azcore.TokenCredential) {
	resourcesClientFactory, err := armresources.NewClientFactory(subscriptionId, credential, nil)
	if err != nil {
		log.Fatalf("Failed to create client factory: %v", err)
	}
	resourceGroupClient = resourcesClientFactory.NewResourceGroupsClient()
}

func getResourceGroup(resourceGroupName string) string {
	resourceGroup, err := resourceGroupClient.Get(context.Background(), resourceGroupName, nil)
	if err != nil {
		log.Fatalf("Failed to retrieve resource group: %+v", err)
	}

	log.Printf("Resource group %s exists", *resourceGroup.ResourceGroup.ID)
	return *resourceGroup.ResourceGroup.ID
}

func createResourceGroup(resourceGroupName string) {
	resourceGroup, err := resourceGroupClient.CreateOrUpdate(context.Background(), resourceGroupName, armresources.ResourceGroup{
		Location: to.Ptr(location),
	}, nil)
	if err != nil {
		log.Fatalf("Creation of resource group failed: %+v", err)
	}

	log.Printf("Resource group %s created", *resourceGroup.ResourceGroup.ID)
}

func deleteResourceGroup(resourceGroupName string) {
	pollerResp, err := resourceGroupClient.BeginDelete(context.Background(), resourceGroupName, nil)
	if err != nil {
		log.Fatalf("Failed to delete resource group: %v", err)
	}

	_, err = pollerResp.PollUntilDone(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get the deployment result: %v", err)
	}
	log.Printf("Resource group %s deleted", resourceGroupName)
}
