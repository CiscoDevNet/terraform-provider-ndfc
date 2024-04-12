#!/bin/bash
set -ex

PATTERN=${1:-TestAcc}
echo $PATTERN
# Set the necessary environment variables
export TF_ACC=1
export TF_LOG=DEBUG
export TF_ACC_LOG=DEBUG
export TF_ACC_LOG_PATH=/tmp/terraform-acceptance-tests.log


# Run the Terraform acceptance tests
rm -rf "$TF_ACC_LOG_PATH"
go test -timeout 1h -v -run ^${PATTERN} ./... 

