#!/bin/sh

# Load variables from .env file
if [ -f .env ]; then
  # Source the .env file
  . .env
else
  echo ".env file not found!"
  exit 1
fi

migrate -path database/migration/ -database "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=require" -verbose up