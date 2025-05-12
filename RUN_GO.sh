#!/bin/bash
source .env

export PSQL=$PSQL
export SERVER_PORT=$SERVER_PORT
export JWT_SECRET=$JWT_SECRET

./app
