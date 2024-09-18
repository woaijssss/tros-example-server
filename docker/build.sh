#!/bin/bash

IMAGE_PREFIX=app
TARGET_NAME=tros_example_server
IMAGE_TAG=`date +'%Y%m%d%H%M'`-`git rev-parse --short HEAD`
img=$IMAGE_PREFIX/${TARGET_NAME}:${IMAGE_TAG}

## run docker
container_id=`docker ps -a | grep $TARGET_NAME | awk '{print $1}'`
if [ ! -z $container_id ];then
  docker stop $container_id
  docker rm $container_id
  docker container prune -f
  docker image prune -a -f
  docker system prune -a -f
fi

go mod vendor
docker build --no-cache -t $img . -f ./docker/'local'.Dockerfile

docker run -d --restart=always --name $TARGET_NAME --net=host -p 9091:9091 -p 9092:9092 -v /etc/localtime:/etc/localtime -v /mnt:/mnt -v /data/logs/$TARGET_NAME/runtime:/$TARGET_NAME/runtime --privileged=true $img
#docker run -d --name trlink_bg --network test_network -p 9091:9091 -p 9092:9092 -v /mnt:/mnt -v /data/logs/trlink_bg/runtime:/trlink_bg/runtime --privileged=true app/trlink_bg:202402041501-7cce817 
cd ..

# 清理无用的资源
rm -rf $TARGET_NAME
docker container prune -f
docker image prune -a -f
docker system prune -a -f
