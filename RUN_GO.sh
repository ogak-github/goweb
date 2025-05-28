#!/bin/bash
source .env

export DATABASE_URL=$PSQL_DEV
export SERVER_PORT=$SERVER_PORT
export JWT_SECRET=$JWT_SECRET

go run main.go
