#!/bin/bash

GOOS=linux GOARCH=amd64 go build -x -o /tmp/web ./cmd/web/
