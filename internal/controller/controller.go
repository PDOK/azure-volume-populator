package controller

import (
	populatorMachinery "github.com/kubernetes-csi/lib-volume-populator/populator-machinery"
	"github.com/pdok/azure-volume-populator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	mountPath  = "/mnt"
	devicePath = "/dev/block"

	groupName  = "volume.pdok.nl"
	apiVersion = "v1alpha1"
	kind       = "AzureVolumePopulator"
	resource   = "azurevolumepopulators"
)

func Run(masterURL string, kubeconfig string, imageName string,
	httpEndpoint string, metricsPath string, namespace string) {

	gk := schema.GroupKind{Group: groupName, Kind: kind}
	gvr := schema.GroupVersionResource{Group: groupName, Version: apiVersion, Resource: resource}
	populatorMachinery.RunController(masterURL, kubeconfig, imageName, httpEndpoint, metricsPath,
		namespace, groupName, gk, gvr, mountPath, devicePath, getPopulatorPodArgs)
}

func getPopulatorPodArgs(rawBlock bool, u *unstructured.Unstructured) ([]string, error) {
	var azure v1alpha1.AzureVolumePopulator
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &azure)
	if nil != err {
		return nil, err
	}
	args := []string{"--mode=populate"}
	if rawBlock {
		args = append(args, "--volume-path="+devicePath)
	} else {
		args = append(args, "--volume-path="+mountPath+"/"+azure.Spec.VolumePath)
	}
	args = append(args, "--blob-prefix="+azure.Spec.BlobPrefix)
	return args, nil
}
