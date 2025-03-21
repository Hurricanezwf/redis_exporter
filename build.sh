#!/bin/bash

set -e

rm -rf ./bin/
mkdir -p ./bin/
GOOS=linux  GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/redis-exporter-v1.68-0.0.1-linux-amd64
GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./bin/redis-exporter-v1.68-0.0.1-darwin-arm64
