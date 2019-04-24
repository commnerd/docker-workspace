#!/bin/bash

docker run -d --restart always --privileged -v ~/Workspace:/home/commnerd/Workspace -v ~/.ssh:/root/.ssh -v /var/run/docker.sock:/var/run/host.docker.sock -p 80:80 -p 9000:9000 --name workspace commnerd/workspace:$1
