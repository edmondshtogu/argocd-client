#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

WORKSPACE="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"

(
    # To support running this from anywhere, first cd into this directory,
    # and then install with forced module mode on and fully qualified name.
    cd "${WORKSPACE}"
    # shellcheck disable=2046 # printf word-splitting is intentional
    GO111MODULE=on go install "k8s.io/code-generator/cmd/client-gen"
)
# Go installs in $GOBIN if defined, and $GOPATH/bin otherwise
gobin="${GOBIN:-$(go env GOPATH)/bin}"

"${gobin}/client-gen" -v "0" \
  --go-header-file "hack/boilerplate.go.txt" \
  --output-dir "pkg/clientset" \
  --output-pkg "github.com/edmondshtogu/argocd-client/pkg/clientset" \
  --clientset-name "versioned" \
  --apply-configuration-package "" \
  --plural-exceptions "" \
  --input-base "${WORKSPACE}/pkg" \
  --input "apis/v1alpha1"