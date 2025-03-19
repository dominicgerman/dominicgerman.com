#!/bin/bash

CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o "$BUILD_PATH" /home/dgerman/hosts/dominicgerman.com/cmd/web/

sudo mv "$BUILD_PATH" /usr/local/bin/portfolio
