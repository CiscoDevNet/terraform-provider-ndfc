#!/bin/bash
# Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#
# SPDX-License-Identifier: MPL-2.0

set -ex

PATTERN=${1:-TestUT}
echo $PATTERN

# Start mockoon server
echo "****** Starting mockoon api server ******"
mockoon-cli start --data ./mockoon_data.json --port 3001 &
sleep 10

# Run the Terraform acceptance tests
echo "****** Begin go test ******"
go test -coverprofile=coverage.out -timeout 1h -v -run ^${PATTERN} terraform-provider-ndfc/internal/provider/ndfc 

# Stop mockoon server
echo "****** Stopping mockoon server******"
pid=$(pgrep -f mockoon-cli.\*3001) ; echo $pid ; kill -9 $pid

# Open coverage report
echo "****** Getting coverage report ******"
go tool cover -html=coverage.out
