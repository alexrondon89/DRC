#!/bin/bash

NAMECSV="csv-generator-service"
OUTPUT_PATH="${1}"
TYPE="${2}"

if [ "$(docker images -q "$NAMECSV")" ]; then
  echo "$NAMECSV image already exists... skipping building step"
else
  echo "creating $NAMECSV image..."
  docker compose -f ./information-collector-service/docker-compose.yaml build --no-cache csv-generator
fi

if [ -n "$(docker ps -q -f name="$NAMECSV")" ]; then
  echo "$NAMECSV container is already running..."
  exit 0
else
  echo "starting $NAMECSV...."
  OUTPUT_PATH="$OUTPUT_PATH" TYPE="$TYPE" docker-compose -f ./information-collector-service/docker-compose.yaml run --rm csv-generator
fi
