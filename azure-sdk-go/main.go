package main

// Import key modules.
import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Define key global variables.
var (
	subscriptionId    = ""
	location          = "swedencentral"
	resourceGroupName = "myResourceGroup" // !! IMPORTANT: Change this to a unique name in your subscription.
	ctx               = context.Background()
)

// Define the function to create a resource group.
func createResourceGroup(subscriptionId string, credential azcore.TokenCredential) (armresources.ResourceGroupsClientCreateOrUpdateResponse, error) {
	rgClient, _ := armresources.NewResourceGroupsClient(subscriptionId, credential, nil)

	param := armresources.ResourceGroup{
		Location: to.Ptr(location),
	}

	return rgClient.CreateOrUpdate(ctx, resourceGroupName, param, nil)
}

// Define the standard 'main' function for an app that is called from the command line.
func main() {

	// Create a credentials object.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
	}

	// Call your function to create an Azure resource group.
	resourceGroup, err := createResourceGroup(subscriptionId, cred)
	if err != nil {
		log.Fatalf("Creation of resource group failed: %+v", err)
	}

	// Print the name of the new resource group.
	log.Printf("Resource group %s created", *resourceGroup.ResourceGroup.ID)
}
