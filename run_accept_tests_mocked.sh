#!/bin/bash
set -x

PATTERN=${1:-TestAcc}
PORT=${2:-3001}
echo $PATTERN
# Set the necessary environment variables
export TF_ACC=1
export TF_LOG=DEBUG
export TF_ACC_LOG=DEBUG
export TF_ACC_LOG_PATH=$(pwd)/terraform-acceptance-tests.log
export NDFC_TEST_CONFIG_FILE=$(pwd)/internal/provider/testing/ndfc_config_mocked.yaml

export NDFC_MOCKED_SERVER="http://localhost:$PORT"
# Run the Terraform acceptance tests
rm -rf "$TF_ACC_LOG_PATH"
go test -timeout 1h -v -run ^${PATTERN} ./... 



