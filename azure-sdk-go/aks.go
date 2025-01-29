package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

var (
	managedClustersClient *armcontainerservice.ManagedClustersClient
)

func createContainerServiceClient(subscriptionId string, credential azcore.TokenCredential) {
	containerserviceClientFactory, err := armcontainerservice.NewClientFactory(subscriptionId, credential, nil)
	if err != nil {
		log.Fatalf("failed to create container service factory: %v", err)
	}
	managedClustersClient = containerserviceClientFactory.NewManagedClustersClient()
}

func createAKSCluster(resourceGroupName string) {

	aksName := fmt.Sprint("aks", time.Now().Unix())

	managedCluster := getManagedCluster(aksName)

	log.Printf("Creating AKS cluster %s", aksName)
	pollerResp, err := managedClustersClient.BeginCreateOrUpdate(
		context.Background(),
		resourceGroupName,
		aksName,
		*managedCluster,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to deploy template: %v", err)
	}

	deploymentResp, err := pollerResp.PollUntilDone(context.Background(), nil)
	if err != nil {
		log.Fatalf("failed to get the deployment result: %v", err)
	}

	log.Printf("AKS cluster %s created", *deploymentResp.ID)
}

func getManagedCluster(aksName string) *armcontainerservice.ManagedCluster {

	managedCluster := &armcontainerservice.ManagedCluster{
		Name:     to.Ptr(aksName),
		Location: to.Ptr(location),
		Identity: &armcontainerservice.ManagedClusterIdentity{
			Type: to.Ptr(armcontainerservice.ResourceIdentityTypeSystemAssigned),
		},
		Properties: &armcontainerservice.ManagedClusterProperties{
			DNSPrefix:         to.Ptr("aksgosdk"),
			KubernetesVersion: to.Ptr("1.30.6"),
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					Name:              to.Ptr("askagent"),
					Count:             to.Ptr[int32](1),
					VMSize:            to.Ptr("Standard_DS2_v2"),
					MaxPods:           to.Ptr[int32](110),
					MinCount:          to.Ptr[int32](1),
					MaxCount:          to.Ptr[int32](3),
					OSType:            to.Ptr(armcontainerservice.OSTypeLinux),
					Type:              to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
					EnableAutoScaling: to.Ptr(true),
					Mode:              to.Ptr(armcontainerservice.AgentPoolModeSystem),
				},
			},
			AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{
				"httpApplicationRouting": {
					Enabled: to.Ptr(true),
				},
			},
		},
	}

	return managedCluster
}
