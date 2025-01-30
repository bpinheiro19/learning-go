package main

import (
	"os"
	"testing"
)

func TestCreateAKSCluster(t *testing.T) {
	subscriptionId := os.Getenv("AZURE_SUBSCRIPTION_ID")
	resourceGroupName := "aks-rg"
	credential := getAzureCredential()

	expectedResult := "/subscriptions/" + subscriptionId + "/resourceGroups/" + resourceGroupName

	createResourceGroupClient(subscriptionId, credential)
	createResourceGroup(resourceGroupName)

	result := getResourceGroup(resourceGroupName)

	if result != expectedResult {
		t.Errorf("Result was incorrect, got: %s, wanted: %s.", result, expectedResult)
	}
}
