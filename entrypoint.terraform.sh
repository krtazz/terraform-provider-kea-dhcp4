#!/bin/sh


sed -i s/localhost/api/ /tffiles/dhcp-provider.tf
terraform init 2>&1 && \
terraform apply -auto-approve 2>&1 && \
terraform destroy -auto-approve 2>&1
