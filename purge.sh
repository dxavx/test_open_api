#!/usr/bin/env bash

docker-compose down --remove-orphans
docker image rm test-open-api-api
docker image rm test-open-api-varnish