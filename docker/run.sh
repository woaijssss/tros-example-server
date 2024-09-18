#!/bin/bash

app=tros_example_server
image_id=app/$app:$1
container_id=`docker ps -a | grep $app | awk '{print $1}'`
docker stop $container_id
docker rm $container_id

docker run -d --name $app --net=host -p 9091:9091 -p 9092:9092 -v /etc/localtime:/etc/localtime -v /mnt:/mnt -v /data/logs/$app/runtime:/$app/runtime --privileged=true $image_id
