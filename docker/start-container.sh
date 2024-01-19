#!/bin/bash
echo "Starting compose files... "

docker compose -f docker-compose-kafka.yml up -d && 
docker compose -f docker-compose-monitoring.yml up -d &&
docker compose -f docker-compose-usermanagement.yml up -d &&
docker compose -f docker-compose-travelmanagement.yml up -d &&
docker compose -f docker-compose-checkout.yml up -d

echo "Starting proxy... "
docker compose -f docker-compose-proxy.yml up -d

echo "Compose files started!"