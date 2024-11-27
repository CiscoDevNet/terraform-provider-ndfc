#!/bin/bash
# Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#
# SPDX-License-Identifier: MPL-2.0

set -e
export PATH=$PATH:$GOPATH/bin
OUTDIR=${1:-"./internal/provider"}
REPLACE=${2}

mkdir -p $OUTDIR

if [[ ! -f $GOPATH/bin/tfplugingen-framework ]]
then
    echo "Install tfplugingen-framework" 
    exit 1 
fi
if [[ ! -f $GOPATH/bin/generator ]]
then
    echo "Compile and install generator"
    exit 1
fi
if [[ ! -f $GOPATH/bin/addlicense ]]
then
    echo "Compile and install addlicense tool"
    exit 1
fi

mkdir -p ./out
rm -rf ./out/*
# Stage 1: Generate spec Json used for TF plugin generator framework
echo "Stage 1 - Generating provider spec JSON"
$GOPATH/bin/generator spec -in ./generator/defs -out ./out -split

# Stage 2: Use TF plugin generator framrework to generate Resource/Provider/Datasource definitions 
# Use Stage 1 output as input here
echo "Stage 2 - Generating TF provider definitions from spec JSON"
FILES=$(ls ./out)
for filename in $FILES; do
    if [[ $filename == *"provider.json" ]]
    then
        echo "Generate provider code from $filename"
        tfplugingen-framework generate provider --input ./out/$filename --output $OUTDIR/provider
    elif [[ $filename == *"resource.json" ]]
    then
        echo "Generate Resource code from $filename"
        tfplugingen-framework generate resources --input ./out/$filename --output $OUTDIR/resources
    elif [[ $filename == *"datasource.json" ]]
    then
        echo "Generate data-source code from $filename"
        tfplugingen-framework generate data-sources --input ./out/$filename --output $OUTDIR/datasources
    fi
done

#Stage 3: Generate additional encoder/decoder code for the generated definitions
#Generated in the same output folder of Stage 2
echo "Stage 3 - Generating additional codec code"
$GOPATH/bin/generator code -in generator/defs -template generator/templates -out $OUTDIR $REPLACE

# Add license headers to all generated files
$GOPATH/bin/addlicense -c "Cisco Systems, Inc. and its affiliates" -l "mpl" -s  internal/provider/resources/**/*_gen.go
$GOPATH/bin/addlicense -c "Cisco Systems, Inc. and its affiliates" -l "mpl" -s  internal/provider/ndfc/*_gen.go
$GOPATH/bin/addlicense -c "Cisco Systems, Inc. and its affiliates" -l "mpl" -s  internal/provider/resources/**/*_test.go
$GOPATH/bin/addlicense -c "Cisco Systems, Inc. and its affiliates" -l "mpl" -s  internal/provider/provider/**/*_gen.go
$GOPATH/bin/addlicense -c "Cisco Systems, Inc. and its affiliates" -l "mpl" -s  internal/provider/datasources/**/*_gen.go
$GOPATH/bin/addlicense -c "Cisco Systems, Inc. and its affiliates" -l "mpl" -s  internal/provider/datasources/**/*_test.go
$GOPATH/bin/addlicense -c "Cisco Systems, Inc. and its affiliates" -l "mpl" -s  internal/provider/types/
$GOPATH/bin/addlicense -c "Cisco Systems, Inc. and its affiliates" -l "mpl" -s  internal/provider/*_test.go

rm -rf ./out






