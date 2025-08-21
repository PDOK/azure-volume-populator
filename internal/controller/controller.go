package controller

import (
	populatorMachinery "github.com/kubernetes-csi/lib-volume-populator/populator-machinery"
	"github.com/pdok/azure-volume-populator/api"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	prefix     = "hello.example.com"
	mountPath  = "/mnt"
	devicePath = "/dev/block"

	groupName  = "hello.example.com"
	apiVersion = "v1alpha1"
	kind       = "Hello"
	resource   = "hellos"
)

func RunController(masterURL string, kubeconfig string, imageName string, httpEndpoint string, metricsPath string,
	namespace string) {
	var (
		gk  = schema.GroupKind{Group: groupName, Kind: kind}
		gvr = schema.GroupVersionResource{Group: groupName, Version: apiVersion, Resource: resource}
	)
	populatorMachinery.RunController(masterURL, kubeconfig, imageName, httpEndpoint, metricsPath,
		namespace, prefix, gk, gvr, mountPath, devicePath, getPopulatorPodArgs)
}

func getPopulatorPodArgs(_ bool, u *unstructured.Unstructured) ([]string, error) {
	var azure api.AzureVolumePopulator
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &azure)
	if nil != err {
		return nil, err
	}
	args := []string{"--mode=populate"}
	args = append(args, "--volume-path="+mountPath+"/"+azure.Spec.VolumePath)
	args = append(args, "--blob-prefix="+azure.Spec.BlobPrefix)
	return args, nil
}
