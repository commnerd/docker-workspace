#!/bin/bash

docker login

./build.sh && docker push commnerd/workspace && docker push commnerd/workspace:php && docker push commnerd/workspace:golang && docker push commnerd/workspace:work && docker push commnerd/jenkins
