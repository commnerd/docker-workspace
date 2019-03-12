#!/bin/bash

docker run -v ~/.go:/go -v ${PWD}/backend:/go/src/github.com/commnerd/docker-workspace/controller/backend -w /go/src/github.com/commnerd/docker-workspace/controller/backend -e  CGO_ENABLED=0 -it --rm --name docker-workspace-controller-builder golang bash -c "go build -tags netgo -a -o workspace-controller"
