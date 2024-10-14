#!/bin/bash

NAME="information-collector-service"
if [ "$(docker images -q "$NAME")" ]; then
  echo "$NAME image already exists... skipping building step"
else
  echo "creating $NAME image..."
  docker compose -f ./information-collector-service/docker-compose.yaml build --no-cache collector
fi

if [ -n "$(docker ps -q -f name="$NAME")" ]; then
  echo "$NAME container is already running..."
  exit 0
else
  echo "starting $NAME...."
  docker-compose -f ./information-collector-service/docker-compose.yaml run --rm collector
fi