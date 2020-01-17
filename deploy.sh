#!/bin/bash

cd ${HOME}/go/src/repo.mlis.one.pl/nutcracker/terraform-provider-kea-dhcp4

go build

if [[ $? == 0 ]]
then
    cp terraform-provider-kea-dhcp4 /home/nutcracker/.terraform.d/plugins/

    cd ${HOME}/terraform-kvm-vmbuild

    source source_this_file.sh

    terraform init

    echo "Do you want to run \"terraform plan\" ? y/n"
    read ANSWER
    if [[ $ANSWER = 'y' ]]
    then
        terraform plan
    fi
    if [[ $? == 0 ]] 
    then    
        echo "Do you want to run \"terraform apply\" ? y/n"
        read ANSWER2
        if [[ $ANSWER2 = 'y' ]]
        then
            terraform apply -auto-approve
        fi
    fi

else
echo -e "\e[41m Problem with building plugin\e[49m"
fi