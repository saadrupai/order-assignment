#!/bin/bash

echo "Building Docker image..."
sudo docker build -t order .

if [ $? -ne 0 ]; then
    echo "Failed to build Docker image."
    exit 1
fi

echo "Starting services with Compose..."
sudo docker-compose up -d


