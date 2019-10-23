#!/usr/bin/env bash

name=$1
username=$2
password=$3
harbor=$4
namespace=$5
kubectl create -n ${namespace} secret generic ${name} --from-literal=username=${username} --from-literal=password=${password} --from-literal=harbor=${harbor}
if [ $? -ne 0 ];then
    echo "create secret: ${name} fail"
    exit 1
fi

echo "create secret: ${name} ok"