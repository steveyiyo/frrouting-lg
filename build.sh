#!/bin/bash

echo "Starting for build!"

echo "Linux amd64"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/Linux_Intel64_lg
echo "macOS amd64"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/macOS_Intel64_lg
echo "Windows amd64"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/Windows_Intel64_lg.exe