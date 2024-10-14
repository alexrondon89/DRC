#!/bin/bash

DB="postgres_1"
PGADMIN="pg_admin_1"

if [ -n "$(docker ps -q -f name="$DB")" ]; then
  echo "$DB container is already running..."
else
  echo "starting $DB container...."
  docker compose -f ./information-collector-service/docker-compose.yaml up postgres -d
fi

if [ -n "$(docker ps -q -f name="$PGADMIN")" ]; then
  echo "$PGADMIN container is already running..."
else
  echo "starting $PGADMIN container...."
  docker compose -f ./information-collector-service/docker-compose.yaml up pgadmin -d
fi