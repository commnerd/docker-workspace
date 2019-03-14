#!/bin/bash

docker run -d --rm --name ws-builder -v ${PWD}:/app -w /app node bash -c "yarn global add @angular/cli && ng build --watch --poll"
