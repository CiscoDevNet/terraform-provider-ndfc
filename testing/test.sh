#!/bin/bash
set -e

# Copy .tf files in order from test_config folder

mkdir  -p ./logs
FROM=${1:=1}
TESTNAME=${2:-"vrf_test"}

TEST_COUNT=$(ls -1 ./test_configs/$TESTNAME*.tf | wc -l)
echo "Total test files: $TEST_COUNT"

# Rename the copied files to main.tf
TS=$(date +%s)
for ((i=$FROM; i<=$TEST_COUNT; i++))
do
    echo "Test starting $i file test_configs/vrf_test_$i.tf"
    file="./test_configs/"${TESTNAME}"_"$i".tf"
    head -1 $file
    cp "$file" "./main.tf"
    TF_LOG=DEBUG TF_LOG_PATH=./logs/tflog_$TS.log terraform apply  -auto-approve || exit 1
    echo "Test completed $i file test_configs/vrf_test_$i.tf"
done


