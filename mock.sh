#!/bin/bash


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
