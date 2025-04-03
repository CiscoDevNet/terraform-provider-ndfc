#!/bin/bash
# Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#
# SPDX-License-Identifier: MPL-2.0



function start_mockoon_server() {
    DATAFILE=${1:-./mockoon_data.json}
    PORT=${2:-3001}

    echo "****** Starting mockoon api server $DATAFILE $PORT $(pwd)******"
    mockoon-cli start --data $DATAFILE --port $PORT &
    sleep 10
    echo "****** Mockoon api server started ******"
}

function stop_mockoon_server() {
    PORT=${1:-3001}
    echo "****** Stopping mockoon server******"
    pid=$(pgrep -f mockoon-cli.\*${PORT}) ; echo $pid ; kill -9 $pid
}

case "$1" in 
    start)
        start_mockoon_server $2 $3
        exit $?
        ;;
    stop)
        stop_mockoon_server $2
        exit $?
        ;;
    *)
        echo "Usage: $0 {start|stop}"
        exit 1
        ;;
esac
