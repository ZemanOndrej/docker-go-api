#!/bin/bash
PORT=8080
INNER_PORT=8080
while getopts p:ip: option
do
case "${option}"
in
p) PORT=${OPTARG};;
ip) INNER_PORT=${OPTARG};;
esac
done


docker build -t golang-api . -e 
docker run --rm -p ${PORT}:${INNER_PORT} golang-api -e PORT=${INNER_PORT} -d
echo "docker container successfully created"