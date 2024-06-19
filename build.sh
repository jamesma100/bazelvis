#!/bin/sh

mkdir -p ./bin
go build -o ./bin/bazelvis ./cmd/bazelvis/main.go
