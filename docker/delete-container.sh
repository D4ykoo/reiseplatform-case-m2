#!/bin/bash

echo "Taking down compose files... "

docker compose -f docker-compose-kafka.yml down &&
docker compose -f docker-compose-monitoring.yml down &&
docker compose -f docker-compose-usermanagement.yml down &&
docker compose -f docker-compose-travelmanagement.yml down &&
docker compose -f docker-compose-checkout.yml down

echo "Taking down proxy... "
docker compose -f docker-compose-proxy.yml down

echo "Compose files taken down!"