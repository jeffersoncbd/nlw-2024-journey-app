#!/bin/bash

set -o allexport
source .env set

docker compose --env-file .env up -d --build
go generate ./...
printf " \033[0;32m✔\033[0m go generate \n"
