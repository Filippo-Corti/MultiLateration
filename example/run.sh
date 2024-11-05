#!/bin/bash

APP_NAME="multilateration"

# Build the Go application (the binary will be stored in ./bin/)
echo "Building the application..."
go build -o ./bin/$APP_NAME ./cmd/server

# Check if build was successful
if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

# Run the Go application
echo "Running the application..."
./bin/$APP_NAME