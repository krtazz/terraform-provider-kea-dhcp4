#!/bin/bash

mkdir build 2>/dev/null || true
DOCKER_BUILDKIT=1 docker build -t terraform-kea-dhcp4 -f terraform.Dockerfile --target exporter -o build .