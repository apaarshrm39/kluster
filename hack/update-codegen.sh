#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

../vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/apaarshrm39/Kluster/pkg/client \
github.com/apaarshrm39/Kluster/pkg/apis \
apaarshrm.dev:v1alpha1 \
 --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.."  \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt