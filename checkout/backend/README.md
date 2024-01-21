# How to run 
### Configuration
Can be applied in the .env file.

### Prerequisites
Create the network if it does not already exist:
```bash
docker network create checkout
```

An up and running postgresdb. Copy the following command for creating a postgres docker container:

```bash
docker run -p 8094:5432  --name checkout-backend-db --network checkout -e POSTGRES_PASSWORD=password -e POSTGRES_USER=checkout -e POSTGRES_DB=checkout -d postgres
```

An up and running kafka. Therefore the docker compose can be used.
 
## Bare Metal
```bash
cargo run
# or build the project and run binary
cargo build --release
./target/release/checkout-backend
```

## Docker
Build the image:
```bash
docker buildx build -t checkout-backend:latest .
```

**THIS WILL GENERATED SECURITY ISSUES IN YOUR NETWORK**
Run the container:
```bash
docker run --name checkout-backend --network host -p 8084:8084 -d checkout-backend 
```

**CURRENTLY NOT WORKING**
Run the container:
```bash
docker run --name checkout-backend --network checkout -p 8084:8084 -d checkout-backend 
```