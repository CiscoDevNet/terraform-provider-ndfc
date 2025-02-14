#!/bin/bash
# Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#
# SPDX-License-Identifier: MPL-2.0

set -ex

PATTERN=${1:-TestAcc}
echo $PATTERN
# Set the necessary environment variables
export TF_ACC=1
export TF_LOG=DEBUG
export TF_ACC_LOG=DEBUG
export TF_ACC_LOG_PATH=/tmp/terraform-acceptance-tests.log
export NDFC_TEST_CONFIG_FILE=$(pwd)/internal/provider/testing/ndfc_config.yaml

# Run the Terraform acceptance tests
rm -rf "$TF_ACC_LOG_PATH"
GOFLAGS="-count=1" go test -timeout 1h -v -run ^${PATTERN} ./... 

