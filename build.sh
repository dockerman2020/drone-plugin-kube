#!/usr/bin/env bash

set -eou pipefail

rm -rf build/kubano
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/kubano

docker build -t dockerman2002/drone-plugin-kube:0.2.0 -t dockerman2002/drone-plugin-kube:latest build
docker push dockerman2002/drone-plugin-kube:0.2.0
docker push dockerman2002/drone-plugin-kube:latest
