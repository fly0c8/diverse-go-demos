#!/bin/bash
GOOGLE_API_PATH="/home/arnold/spark/project/GoCode/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.4/third_party/googleapis"
echo $GOOGLE_API_PATH
#protoc --go_out=plugins=grpc:chat chat.proto
#protoc -I$GOOGLE_API_PATH --swagger_out=logtostderr=true:chat --proto_path . chat.proto
protoc -I$GOOGLE_API_PATH \
    --go_out=plugins=grpc:chat \
    --grpc-gateway_out=logtostderr=true:chat \
    --swagger_out=logtostderr=true:chat \
    --proto_path . chat.proto

