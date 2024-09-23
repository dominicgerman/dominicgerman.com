#!/bin/bash

CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o /usr/local/bin/devblog ./cmd/web/
