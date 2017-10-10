#!/usr/bin/env bash
set -e
GOOS=linux go build

docker build -t lbaker/testserver .
docker push lbaker/testserver
go clean