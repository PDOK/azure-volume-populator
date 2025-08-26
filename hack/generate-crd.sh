#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
cd "${REPO_ROOT}"

if ! command -v $(pwd)/bin/controller-gen &> /dev/null
then
  GOBIN=$(pwd)/bin go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.18.0
fi

$(pwd)/bin/controller-gen object paths="$(pwd)/api/..." output:dir="$(pwd)/api/v1alpha1"
$(pwd)/bin/controller-gen crd paths="$(pwd)/api/..." output:dir="$(pwd)/config/crd"

# Assertions
cat config/crd/volume.pdok.nl_azurevolumepopulators.yaml | grep "kind: AzureVolumePopulator"

