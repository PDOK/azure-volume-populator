package main

import (
	"os"

	"github.com/pdok/azure-volume-populator/internal/controller"
	"github.com/pdok/azure-volume-populator/internal/populator"
	"github.com/urfave/cli/v2"
	"k8s.io/klog/v2"
)

var (
	cliFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     "mode",
			Usage:    "Mode to run in (controller, populate)",
			EnvVars:  []string{"MODE"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     "azure-storage-connection-string",
			Usage:    "connection string to access data in an Azure Storage Account.",
			EnvVars:  []string{"AZURE_STORAGE_CONNECTION_STRING"},
			Required: true,
		},
		// Populate args
		&cli.StringFlag{
			Name:  "blob-prefix",
			Usage: "Copy all Azure blobs with this prefix (can be multiple files). Should take the form of a container name + path within container e.g. 'mycontainer/firstfolder/secondfolder/etc'.",
		},
		&cli.UintFlag{
			Name:  "blob-block-size",
			Usage: "Block size to use when downloading from Azure Blob Storage.",
			Value: 4 * 1024 * 1024,
		},
		&cli.UintFlag{
			Name:  "blob-concurrency",
			Usage: "Number of blobs to download concurrently.",
			Value: 10,
		},
		&cli.StringFlag{
			Name:  "volume-path",
			Usage: "Destination path on the volume.",
		},
		// Controller args
		&cli.StringFlag{
			Name:  "kubeconfig",
			Usage: "Path to a kubeconfig. Only required if out-of-cluster.",
		},
		&cli.StringFlag{
			Name:  "master",
			Usage: "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.",
		},
		&cli.StringFlag{
			Name:  "image-name",
			Usage: "Image to use for populating.",
		},
		&cli.StringFlag{
			Name:  "namespace",
			Usage: "Namespace to deploy controller.",
		},
		// Metrics args
		&cli.StringFlag{
			Name:  "http-endpoint",
			Usage: "The TCP network address where the HTTP server for diagnostics, including metrics and leader election health check, will listen (example: `:8080`). The default is empty string, which means the server is disabled.",
		},
		&cli.StringFlag{
			Name:  "metrics-path",
			Usage: "The HTTP path where prometheus metrics will be exposed. Default is `/metrics`.",
			Value: "/metrics",
		},
	}
)

func main() {
	app := &cli.App{
		Name:  "azure-volume-populator",
		Usage: "Run either as Kubernetes controller or in populate mode",
		Flags: cliFlags,
		Action: func(ctx *cli.Context) error {
			mode := ctx.String("mode")
			azConnectionString := ctx.String("azure-storage-connection-string")
			blobPrefix := ctx.String("blob-prefix")
			blobBlockSize := ctx.Uint("blob-block-size")
			blobConcurrency := ctx.Uint("blob-concurrency")
			volumePath := ctx.String("volume-path")
			kubeconfig := ctx.String("kubeconfig")
			masterURL := ctx.String("master")
			imageName := ctx.String("image-name")
			namespace := ctx.String("namespace")
			httpEndpoint := ctx.String("http-endpoint")
			metricsPath := ctx.String("metrics-path")

			switch mode {
			case "controller":
				controller.Run(masterURL, kubeconfig, imageName, httpEndpoint, metricsPath, namespace, azConnectionString)
			case "populate":
				populator.Populate(blobPrefix, volumePath, blobBlockSize, blobConcurrency, azConnectionString)
			default:
				klog.Fatalf("invalid mode: %s", mode)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		klog.Fatalf("application error: %v", err)
	}
}
