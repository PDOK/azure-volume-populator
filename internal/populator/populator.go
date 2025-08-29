package populator

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"k8s.io/klog/v2"
)

func Populate(blobPrefix, volumePath string, blockSize, concurrency uint, azConnectionString string) {
	if azConnectionString == "" {
		klog.Fatalf("Missing required arg --azure-storage-connection-string")
	}
	if blobPrefix == "" {
		klog.Fatalf("Missing required arg --blob-prefix")
	}
	if volumePath == "" {
		klog.Fatalf("Missing required arg --volume-path")
	}

	client, err := azblob.NewClientFromConnectionString(azConnectionString, nil)
	if err != nil {
		klog.Fatalf("Failed to create client: %v", err)
	}

	containerName, blobPrefixWithoutContainerName := getContainerNameAndPath(blobPrefix)

	pager := client.NewListBlobsFlatPager(containerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{
			Snapshots: false,
			Versions:  false,
		},
		Prefix: &blobPrefixWithoutContainerName,
	})

	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			klog.Fatalf("Failed to list blobs in container '%s' with prefix '%s': %v",
				containerName, blobPrefixWithoutContainerName, err)
		}

		for _, blob := range page.Segment.BlobItems {
			klog.Infof("Downloading blob '%s' from container '%s' (with block size %d, concurrency %d)",
				*blob.Name, containerName, blockSize, concurrency)

			// Create a directory structure on volume matching the blob path
			filePath := filepath.Join(volumePath, *blob.Name)
			if err = os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
				klog.Fatalf("Download of blob '%s' failed: cannot create directory: %v", *blob.Name, err)
			}

			// Create a file on volume to hold blob data
			file, err := os.Create(filePath)
			if err != nil {
				klog.Fatalf("Download of blob '%s' failed: cannot create file: %v", *blob.Name, err)
			}
			if err = file.Close(); err != nil {
				klog.Fatalf("Download of blob '%s' failed: cannot close file: %v", *blob.Name, err)
			}

			// Download blob to file
			if _, err = client.DownloadFile(context.Background(), containerName, *blob.Name, file, &azblob.DownloadFileOptions{
				BlockSize:   int64(blockSize),    //nolint:gosec // no risk of overflow
				Concurrency: uint16(concurrency), //nolint:gosec // no risk of overflow
			}); err != nil {
				klog.Fatalf("Download of blob '%s' failed during transfer: %v", *blob.Name, err)
			}
		}
	}
}

func getContainerNameAndPath(blobPrefix string) (containerName string, blobPath string) {
	if !strings.HasSuffix(blobPrefix, "/") {
		blobPrefix += "/"
	}
	parts := strings.SplitN(blobPrefix, "/", 2)
	containerName = parts[0]
	blobPath = parts[1]
	return
}
