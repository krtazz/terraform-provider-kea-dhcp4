#!/bin/bash

DOCKER_BUILDKIT=1 docker build -t terraform-kea-dhcp4 -f terraform.Dockerfile --target exporter -o build .
