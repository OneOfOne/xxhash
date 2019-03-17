#!/bin/sh

VERSION=$1

env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o xxhsum64-${VERSION}.exe || exit 1
env GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -ldflags "-s -w" -o xxhsum32-${VERSION}.exe || exit 1
