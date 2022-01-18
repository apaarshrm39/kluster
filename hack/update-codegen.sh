#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail



../vendor/k8s.io/code-generator/generate-groups.sh "deepcopy,client,informer,lister" \
github.com/apaarshrm39/Kluster/pkg/client \
github.com/apaarshrm39/Kluster/pkg/apis \
  apaarshrm.dev:v1alpha1 \
  --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
  --go-header-file ../hack/boilerplate.go.txt