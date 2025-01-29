package main

import (
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

const (
	location       = "swedencentral"
	deploymentName = "deployARM"
)

func getAzureCredential() azcore.TokenCredential {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
	}
	return cred
}

func main() {

	subscriptionId := os.Getenv("AZURE_SUBSCRIPTION_ID")

	if len(subscriptionId) == 0 {
		log.Fatal("AZURE_SUBSCRIPTION_ID is not set.")
	}

	resourceGroupName := "aks-rg"
	credential := getAzureCredential()

	createResourceGroupClient(subscriptionId, credential)
	createContainerServiceClient(subscriptionId, credential)

	createResourceGroup(resourceGroupName)

	createAKSCluster(resourceGroupName)

	deleteResourceGroup(resourceGroupName)
}
