package populator

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"k8s.io/klog/v2"
)

func Populate(blobPrefix, volumePath, azConnectionString string) {
	if "" == blobPrefix || "" == volumePath {
		klog.Fatalf("Missing required arg")
	}
	// TODO implement
	_, err := azblob.NewClientFromConnectionString(azConnectionString, nil)
	if err != nil {
		klog.Fatalf("Failed to create client: %v", err)
	}
}
