#!/bin/bash
# Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#
# SPDX-License-Identifier: MPL-2.0

set -x

PATTERN=${1:-TestAccInterface}
PORT=${2:-3001}
echo $PATTERN
# Set the necessary environment variables
export TF_ACC=1
export TF_LOG=DEBUG
export TF_ACC_LOG=DEBUG
export TF_ACC_LOG_PATH=/tmp/terraform-acceptance-tests.log
export NDFC_TEST_CONFIG_FILE=$(pwd)/internal/provider/testing/ndfc_config.yaml

export NDFC_MOCKED_SERVER="http://localhost:$PORT"
# Run the Terraform acceptance tests
rm -rf "$TF_ACC_LOG_PATH"
go test -coverprofile ./cover.out -timeout 1h -v -run ^${PATTERN} ./... -coverpkg=./...



