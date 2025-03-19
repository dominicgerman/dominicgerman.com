#!/bin/bash

PROJECT_DIR="/home/dgerman/hosts/dominicgerman.com"
cd "$PROJECT_DIR" || { echo "Failed to cd into $PROJECT_DIR"; exit 1; }

CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o "$BUILD_PATH" ./cmd/web/

if [[ ! -f "$BUILD_PATH" ]]; then
    echo "Error: Build output $BUILD_PATH does not exist."
    exit 1
fi

sudo mv "$BUILD_PATH" /usr/local/bin/portfolio
