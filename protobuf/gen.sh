#!/bin/bash

for d in $(find . -mindepth 1 -type d -not -path '*/\.*'); do
    pushd $d &> /dev/null

    protoc -I/usr/local/include -I. \
        -I$GOPATH/src \
        -I${GOPATH}/src/github.com/rnd/kudu/protobuf \
        --go_out=plugins=grpc:$GOPATH/src \
        *.proto

    popd &> /dev/null
done