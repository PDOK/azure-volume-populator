package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pdok/azure-volume-populator/internal/controller"
	"github.com/pdok/azure-volume-populator/internal/populator"
	"k8s.io/klog/v2"
)

var version = "unknown"

func main() {
	var (
		mode         string
		blobPrefix   string
		volumePath   string
		httpEndpoint string
		metricsPath  string
		masterURL    string
		kubeconfig   string
		imageName    string
		showVersion  bool
		namespace    string
	)
	klog.InitFlags(nil)
	// Main arg
	flag.StringVar(&mode, "mode", "", "Mode to run in (controller, populate)")
	// Populate args
	flag.StringVar(&blobPrefix, "blob-prefix", "", "Copy all Azure blobs with this prefix")
	flag.StringVar(&volumePath, "volume-path", "", "Destination path on the volume")
	// Controller args
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&imageName, "image-name", "", "Image to use for populating")
	// Metrics args
	flag.StringVar(&httpEndpoint, "http-endpoint", "", "The TCP network address where the HTTP server for diagnostics, including metrics and leader election health check, will listen (example: `:8080`). The default is empty string, which means the server is disabled.")
	flag.StringVar(&metricsPath, "metrics-path", "/metrics", "The HTTP path where prometheus metrics will be exposed. Default is `/metrics`.")
	// Other args
	flag.BoolVar(&showVersion, "version", false, "display the version string")
	flag.StringVar(&namespace, "namespace", "", "Namespace to deploy controller")
	flag.Parse()

	if showVersion {
		fmt.Println(os.Args[0], version)
		os.Exit(0)
	}

	if mode == "" {
		klog.Fatalf("Missing required arg: --mode")
	}
	switch mode {
	case "controller":
		controller.Run(masterURL, kubeconfig, imageName, httpEndpoint, metricsPath, namespace)
	case "populate":
		populator.Populate(blobPrefix, volumePath)
	default:
		klog.Fatalf("Invalid mode: %s", mode)
	}
}
