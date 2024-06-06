#!/bin/bash

SERVICE_NAME=$1
RELEASE_VERSION=$2
USER_NAME=$3
EMAIL=$4

echo "SERVICE_NAME: $SERVICE_NAME"
echo "RELEASE_VERSION: $RELEASE_VERSION"
echo "USER_NAME: $USER_NAME"
echo "EMAIL: $EMAIL"

git config user.name "$USER_NAME"
git config user.email "$EMAIL"
git fetch --all && git checkout main

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc --go_out=./golang --go_opt=paths=source_relative \
  --go-grpc_out=./golang --go-grpc_opt=paths=source_relative \
 ./${SERVICE_NAME}/*.proto
cd golang/${SERVICE_NAME}
go mod init \
  github.com/ayuved/microservices-proto/golang/${SERVICE_NAME} || true
go mod tidy
cd ../../
echo "git"
git add . && git commit -am "proto update" || true
git push origin HEAD:main
git tag -fa golang/${SERVICE_NAME}/${RELEASE_VERSION} \
  -m "golang/${SERVICE_NAME}/${RELEASE_VERSION}" 
git push origin refs/tags/golang/${SERVICE_NAME}/${RELEASE_VERSION}