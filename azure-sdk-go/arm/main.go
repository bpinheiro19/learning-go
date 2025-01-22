package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

var (
	subscriptionId = os.Getenv("AZURE_SUBSCRIPTION_ID")
	ctx            = context.Background()
)

const (
	location           = "swedencentral"
	resourceGroupName  = "aks-rg"
	deploymentName     = "deployARM"
	templateFileName   = "template.json"
	parametersFileName = "parameters.json"
)

func readJSON(path string) map[string]interface{} {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	contents := make(map[string]interface{})

	err = json.Unmarshal(data, &contents)
	if err != nil {
		log.Fatalf("failed to unmarshal file: %v", err)
	}

	return contents
}

func getAzureCredential() azcore.TokenCredential {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
	}
	return cred
}

func createResourceGroup(credential azcore.TokenCredential) {
	rgClient, err := armresources.NewResourceGroupsClient(subscriptionId, credential, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	param := armresources.ResourceGroup{
		Location: to.Ptr(location),
	}

	resourceGroup, err := rgClient.CreateOrUpdate(context.Background(), resourceGroupName, param, nil)

	if err != nil {
		log.Fatalf("Creation of resource group failed: %+v", err)
	}

	log.Printf("Resource group %s created", *resourceGroup.ResourceGroup.ID)
}

func createAKSCluster(credential azcore.TokenCredential) {

	templateFile := readJSON(templateFileName)

	parametersFile := readJSON(parametersFileName)
	parametersFile["clusterName"] = map[string]string{"value": "myAKSCluster"}

	deploymentsClient, err := armresources.NewDeploymentsClient(subscriptionId, credential, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	pollerResp, err := deploymentsClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		deploymentName,
		armresources.Deployment{
			Properties: &armresources.DeploymentProperties{
				Template:   templateFile,
				Parameters: parametersFile,
				Mode:       to.Ptr(armresources.DeploymentModeIncremental),
			},
		},
		nil,
	)
	if err != nil {
		log.Fatalf("failed to deploy template: %v", err)
	}

	deploymentResp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to get the deployment result: %v", err)
	}

	log.Printf("AKS cluster %s created", *deploymentResp.Name)
}

func main() {

	credential := getAzureCredential()

	createResourceGroup(credential)

	createAKSCluster(credential)

}
