# Checkout/ Payment Microservice
Frontend and backend can be deployed seperatly. For simplicty of this lecture it is one directory in a monorepo instead of two seperate repositories.

## Prerequisites
* A postgresql database where the configuration and credentials are known
* See [backend/README.md](backend/README.md) for docker postgresdb instructions (the network arg can be removed when the other services are not deployed with docker) 
* Make sure to install `npm`, `angular` and `rustc` compared with `cargo` when running in bare metal mode

## Configuration
All configuration instructions can be viewed in the corresponding READMEs of the backend and frontend directory.

## How to run - VM
Follow the instructions in the [RunInVM.md](RunInVM.md) of this directory.

## How to run - Docker
Follow the instructions in the READMEs located at:
[/frontend/README.md](/frontend/README.md) and [/backend/README.md](/backend/README.md).

<!-- #### Run with image from docker hub registry 
```bash
docker compose -f docker-compose.yml up -d 
``` -->
## How to run the k8s manifests
Follow the instructions located in [ApplyManifests.md](ApplyManifests.md)

## How to run - Bare Metal
Follow the instructions in the [RunBareMetal.md](RunBareMetal.md) of this directory.

## How the multi platform images were made
Create the builder:
```bash
 sudo docker buildx create --name armbuilder --driver=docker-container
```
Then build and push the multi platform images: 
```bash
sudo docker buildx build --push --platform linux/amd64,linux/arm64,linux/arm/v7 --builder=armbuilder -t dak4408/travma-usermanagement-<checkout/backend>:latest .
```

## Technology Stack
* Angular + typescript + npm + TailwindCSS + daisyUI
* Rust + Actix Web + Diesel

## Additional Information
* Project management done with self-hosted [Leantime](https://github.com/Leantime/leantime)
* Verification of the JWT: 
  * Username included 
  * -> extract JWT via browser dev tools and paste into [https://jwt.io/](https://jwt.io/)
  * Different users have different usernames

## Future Work
* Add UUIDs for db scheme
* Payment method using actual services 