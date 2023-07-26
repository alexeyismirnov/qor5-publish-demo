#!/bin/sh

docker-compose up -d && source dev_env && go run main.go
