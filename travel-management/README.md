# Travel Management Microservice
Frontend and backend can be deployed seperatly. For simplicty of this lecture it is one directory in a monorepo instead of two seperate repositories.

## Prerequisites
* A postgresql database with known configuration and credentials
* See [backend/README.md](backend/README.md) for docker postgresdb instructions (the network arg can be removed when the other services are not deployed with docker) 
* Make sure to install `npm` and `golang` when running in bare metal mode
## Configuration
All the configuration for the backend is done in the .env file located in [backend/.env](backend/.env) . See comments for some additional information.

The `apiURL` in the frontend: [frontend/src/assets/config.ts](frontend/src/assets/config.ts) 
must match with the config for the `API_URL` line, located in the .env file.

## How to run - VM
Follow the instructions in the [RunInVM.md](RunInVM.md) of this directory.

## How to run - Docker
Follow the instructions in the READMEs located at:
[/frontend/README.md](/frontend/README.md) and [/backend/README.md](/backend/README.md).

## How to run - Docker Compose
An All-in-One Solution is provided in the top level `docker` directory. Follow the [/docker/README.md](../docker/README.md) instructions for more information.

When running the frontend and backend seperatly follow the instructions in the READMEs located at:
[/frontend/README.md](/frontend/README.md) and [/backend/README.md](/backend/README.md).

## How to run - Bare Metal
Follow the instructions in the [RunBareMetal.md](RunBareMetal.md) of this directory.

## How the multi platform images were made
Create the builder:
```bash
 sudo docker buildx create --name armbuilder --driver=docker-container
```
Then build and push the multi platform images: 
```bash
sudo docker buildx build --push --platform linux/amd64,linux/arm64 --builder=armbuilder -t dak4408/travma-travelmanagement-<frontend/backend>:latest .
```

## Technology Stack
* Angular + typescript + npm + ngprime
* Go + Gin + GORM

## Architecture
![architecture](./resources/travelmanagement_architecture.png?raw=true "Travelmanagement Architecture")

## Future Work
* Delete Travels after checkout
* Add more filter possibilities
* Block Buy Option, if offer is in a card