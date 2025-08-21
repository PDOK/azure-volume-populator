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

func getPopulatorPodArgs(rawBlock bool, u *unstructured.Unstructured) ([]string, error) {
	var hello api.Hello
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &hello)
	if nil != err {
		return nil, err
	}
	args := []string{"--mode=populate"}
	if rawBlock {
		args = append(args, "--file-name="+devicePath)
	} else {
		args = append(args, "--file-name="+mountPath+"/"+hello.Spec.FileName)
	}
	args = append(args, "--file-contents="+hello.Spec.FileContents)
	return args, nil
}
