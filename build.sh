#!/bin/bash
set -e
DIR=$(cd "$(dirname $0)" && pwd)
APP_NAME="${DIR##*/}"
cd $DIR
rm -rf target
mkdir -p target/bin
mkdir -p target/conf
cp conf/seezoon.yml target/conf

#\cp -R build/assembly/* build/package/$APP_NAME
# -o 后加目录则放入目录，不是目录则为产出物名称
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o target/bin/$APP_NAME cmd/main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/bin/$APP_NAME cmd/main.go

docker build --build-arg MODULE_NAME=${APP_NAME} -t "${APP_NAME}" .