#!/usr/bin/env bash
set -e
GOOS=linux go build

docker build -t lbaker/zipsvr .
docker push lbaker/zipsvr
go clean