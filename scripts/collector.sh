#!/bin/bash

NAME="information-collector-service"
if [ "$(docker images -q "$NAME")" ]; then
  echo "$NAME image already exists... skipping building step"
else
  echo "creating $NAME image..."
  docker build -t "$NAME":latest -f ./information-collector-service/Dockerfile ./information-collector-service/
fi

until docker exec postgres_1 pg_isready -U postgres; do
  echo "waiting for postgres..."
  sleep 2
done

if [ -n "$(docker ps -q -f name="$NAME")" ]; then
  echo "$NAME container is already running..."
  exit 0
else
  echo "starting $NAME...."
  docker run --rm "$NAME"
fi