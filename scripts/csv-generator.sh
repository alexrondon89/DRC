#!/bin/bash

NAME="csv-generator-service"
if [ "$(docker images -q "$NAME")" ]; then
  echo "$NAME image already exists... skipping building step"
else
  echo "creating $NAME image..."
  docker build --no-cache -t "$NAME":latest -f ./csv-generator-service/Dockerfile ./csv-generator-service/
fi

until docker exec postgres_1 pg_isready -U postgres; do
  echo "waiting for postgres..."
  sleep 2
done

if [ -n "$(docker ps -q -f name="$NAME")" ]; then
  echo "$NAME container is already running..."
  exit 0
else
  echo "path to save the file ${1}"
  docker run --rm -v "${1}":./ "$NAME"
fi
