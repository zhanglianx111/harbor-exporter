#!/usr/bin/env bash
if [ -f harbor-exporter ];then
    rm -f harbor-exporter
fi
CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build
docker build -t harbor.aibee.cn/platform/harbor-exporter:dev .
docker push harbor.aibee.cn/platform/harbor-exporter:dev

rm harbor-exporter