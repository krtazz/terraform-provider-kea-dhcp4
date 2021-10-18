#!/bin/bash


set -e

docker kill kea api 2>/dev/null || true
docker rm kea api terraform 2>/dev/null || true
docker network rm kea 2>/dev/null || true
docker network create kea
docker build -t kea-test -f kea.Dockerfile .
docker build -t nginx-test -f api.Dockerfile .
docker build -t terraform-kea-dhcp4 -f terraform.Dockerfile . 
docker run -d --network kea --name kea kea-test 
docker run -d --network kea --name api -p 127.0.0.1:8080:8080/tcp nginx-test
docker run --network kea --name terraform terraform-kea-dhcp4
docker kill kea api 2>/dev/null || true
docker network rm kea
