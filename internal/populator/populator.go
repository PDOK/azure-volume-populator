package populator

import (
	"k8s.io/klog/v2"
)

func Populate(blobPrefix, volumePath string) {
	if "" == blobPrefix || "" == volumePath {
		klog.Fatalf("Missing required arg")
	}
	// TODO implement
}
