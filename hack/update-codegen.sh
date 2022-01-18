#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

package_path="/Users/yuanmh/Desktop/code/k8s-crds-clientsets"


../vendor/k8s.io/code-generator/generate-groups.sh "deepcopy,client,informer,lister" \
../pkg/client \
../pkg/apis \
--output-base "../../../.."  \
  apaarshrm.dev:v1alpha1 \
  --go-header-file ../hack/boilerplate.go.txt