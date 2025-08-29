# Azure Volume Populator

_Kubernetes [Volume Populator](https://kubernetes.io/blog/2025/05/08/kubernetes-v1-33-volume-populators-ga/) for [Azure Blob Storage](https://azure.microsoft.com/en-us/products/storage/blobs/)._

[![Build](https://github.com/PDOK/azure-volume-populator/actions/workflows/build-and-publish-image.yml/badge.svg)](https://github.com/PDOK/azure-volume-populator/actions/workflows/build-and-publish-image.yml)
[![Lint](https://github.com/PDOK/azure-volume-populator/actions/workflows/lint-go.yml/badge.svg)](https://github.com/PDOK/azure-volume-populator/actions/workflows/lint-go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pdok/azure-volume-populator)](https://goreportcard.com/report/github.com/pdok/azure-volume-populator)
[![GitHub license](https://img.shields.io/github/license/PDOK/azure-volume-populator)](https://github.com/PDOK/gokoala/blob/master/LICENSE)
[![Docker Pulls](https://img.shields.io/docker/pulls/pdok/azure-volume-populator.svg)](https://hub.docker.com/r/pdok/azure-volume-populator)

This application supports two modes:
- `controller`: Run as a Kubernetes controller that watches a [AzureVolumePopulator](config/crd/volume.pdok.nl_azurevolumepopulators.yaml) custom resource and schedules a pod to populate volumes using this same binary in populate mode.
- `populate`: Run as a standalone tool, started by the mentioned Kubernetes controller, to actually populate a PersistentVolume with blobs from Azure Blob Storage.

It supports downloading all blobs under a given prefix (container/path).

## Build

Using Docker:

```bash
docker build -t pdok/azure-volume-populator .
```

Using Go:

```bash
go mod download
go generate ./...
go build -o azure-volume-populator cmd/main.go
```

## Run

In controller mode:

```bash
export MODE=controller
export AZURE_STORAGE_CONNECTION_STRING="DefaultEndpointsProtocol=...;AccountName=...;AccountKey=...;EndpointSuffix=core.windows.net"

./bin/azure-volume-populator \
  --mode=controller \
  --azure-storage-connection-string="$AZURE_STORAGE_CONNECTION_STRING" \
  --image-name="docker.io/pdok/azure-volume-populator:latest" \
  --namespace="default" \
  --http-endpoint=":8080" \
  --metrics-path="/metrics"
```

In populate mode:

```bash
export MODE=populate
export AZURE_STORAGE_CONNECTION_STRING="DefaultEndpointsProtocol=...;AccountName=...;AccountKey=...;EndpointSuffix=core.windows.net"

./bin/azure-volume-populator \
  --mode=populate \
  --azure-storage-connection-string="$AZURE_STORAGE_CONNECTION_STRING" \
  --blob-prefix="mycontainer/folder/subfolder" \
  --volume-path="/data" \
  --blob-block-size=4194304 \
  --blob-concurrency=10
```

## Misc

### How to Contribute

Make a pull request...

### Contact

Contacting the maintainers can be done through the issue tracker.