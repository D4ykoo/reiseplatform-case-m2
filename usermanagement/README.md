# User Management
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
#### Run with local (registry) built image
```bash
docker compose -f docker-compose-local.ylm up -d 
```

#### Run with image from docker hub registry 
```bash
docker compose -f docker-compose.ylm up -d 
```

## How to run - Bare Metal
Follow the instructions in the [RunBareMetal.md](RunBareMetal.md) of this directory.

## How the multi platform images were made
Create the builder:
```bash
 sudo docker buildx create --name armbuilder --driver=docker-container
```
Then build and push the multi platform images: 
```bash
sudo docker buildx build --push --platform linux/amd64,linux/arm64,linux/arm/v7 --builder=armbuilder -t dak4408/travma-usermanagement-<frontend/backend>:latest .
```

## Technology Stack
* TypeScript + VueJS 3 + TailwindCSS + daisyUI 
* Go + Gin + GORM

## Architecture
![architecture](./resources/usermangement_architecture.png?raw=true "Usermanagement Architecture")

## Project Structure
The project is separated in frontend and backend directories as well as resources for all the documentation resources.<br>
The following represents a quick and small overview about the microservice structure.
```bash
├── backend
│   ├── adapter
│   ├── application
│   ├── domain
│   │   └── model
│   ├── ports
│   ├── tests
│   ├── utils
│   ├── .env
│   ├── README.md
│   ├── Dockerfile
│
├── frontend
│   ├── public
│   ├── src
│   │   └── assets
│   │   └── components
│   │   └── models
│   │   └── router
│   │   └── services
│   │   └── store
│   │   └── views
│   ├── README.md
│   ├── Dockerfile
│ 
├── resources

```

## Additional Information
* Project management done with self-hosted [Leantime](https://github.com/Leantime/leantime)
* Verification of the JWT: 
  * Username included 
  * -> extract JWT via browser dev tools and paste into [https://jwt.io/](https://jwt.io/)
  * Different users have different usernames

## Future Work
* Add UUIDs for db scheme
* Make username and email unique in DB
* Some more enhanced frontend configuration 
* Reset password link generation
* SMTP for email 