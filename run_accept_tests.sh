#!/bin/bash
# Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#
# SPDX-License-Identifier: MPL-2.0

set -ex
echo "export ACC_TEST_CFG to use a different test environment configuration file; default testing/ndfc_config.yaml"
PATTERN=${1:-TestAcc}
CFG=${TESTBED:-175}
TIMEOUT=${TEST_TIMEOUT:-5h}
echo $PATTERN
# Set the necessary environment variables
export TF_ACC=1
export TF_LOG=DEBUG
export TF_ACC_LOG=DEBUG
export TF_ACC_LOG_PATH=/tmp/terraform-acceptance-tests.log
export NDFC_TEST_CONFIG_FILE=$(pwd)/testing/at_testbeds/ndfc_${CFG}.yaml

# Run the Terraform acceptance tests
rm -rf "$TF_ACC_LOG_PATH"
GOFLAGS="-count=1" go test -timeout ${TIMEOUT} -v -run ^${PATTERN} ./... 

