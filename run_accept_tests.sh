#!/bin/bash
# Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#
# SPDX-License-Identifier: MPL-2.0

set -ex
echo "export TESTBED_FILE to use a different test environment configuration file; default testing/testbed.yaml"
PATTERN=${1:-TestAcc}
CWD=$(pwd)
TIMEOUT=${TEST_TIMEOUT:-5h}
TEST_CFG=${TESTBED_FILE:-$(pwd)/testing/testbed.yaml}
echo $PATTERN
# Set the necessary environment variables
export TF_ACC=1
export TF_LOG=DEBUG
export TF_ACC_LOG=DEBUG
TIME_DATE=$(date '+%Y-%m-%d-%H-%M-%S')
export TF_ACC_LOG_PATH=/tmp/terraform-acceptance-tests_${TIME_DATE}.log
export NDFC_TEST_CONFIG_FILE=${TEST_CFG}

# Run the Terraform acceptance tests
rm -rf "$TF_ACC_LOG_PATH"
# install test summary tool
go get gotest.tools/gotestsum
go install gotest.tools/gotestsum
#GOFLAGS="-count=1" go test -timeout ${TIMEOUT} -v -run ^${PATTERN} ./... 
GOFLAGS="-count=1" $GOBIN/gotestsum --format testname  --format-hide-empty-pkg  --debug --jsonfile /tmp/tftest_output_${TIME_DATE}.json -- -failfast -v -timeout ${TIMEOUT} -run ^${PATTERN} ./internal/provider
