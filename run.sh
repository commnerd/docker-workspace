#!/bin/bash

docker run -d --restart always --privileged -v ~/Workspace:/home/commnerd/Workspace -v ~/.ssh:/root/.ssh -v /var/run/docker.sock:/var/run/host.docker.sock --network host --name workspace commnerd/workspace
