#!/bin/sh

go mod tidy
make lint-full
go test ./...