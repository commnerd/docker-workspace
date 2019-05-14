#!/bin/bash

docker build --squash -t commnerd/workspace . || docker build -t commnerd/workspace .
docker build --squash -t commnerd/workspace:php -f Dockerfile.php . || docker build -t commnerd/workspace:php -f Dockerfile.php .
docker build --squash -t commnerd/workspace:golang -f Dockerfile.go . || docker build -t commnerd/workspace:golang -f Dockerfile.go .
docker build --squash -t commnerd/workspace:work -f Dockerfile.work . || docker build -t commnerd/workspace:work -f Dockerfile.work .
docker build --squash -t commnerd/jenkins -f Dockerfile.jenkins . || docker build -t commnerd/jenkins -f Dockerfile.jenkins .

docker system prune --all
