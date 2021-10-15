#!/bin/sh

terraform init 2>&1 && \
terraform apply -auto-approve 2>&1 && \
terraform destroy -auto-approve 2>&1

echo "Press CTRL-C to tear down test environment"
