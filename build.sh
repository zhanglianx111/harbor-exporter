#!/usr/bin/env bash

if [ $# -eq 0 ];then
    echo "usage:
    $0 tag
    "
    exit 0
fi

tag=$1

docker build -t harbor.aibee.cn/platform/harbor-exporter:$tag --no-cache .
docker push harbor.aibee.cn/platform/harbor-exporter:$tag
