# How to run 
### Configuration
Can be applied in the .env file.

### Prerequisites
Create the network if it does not already exist:
```bash
docker network create monitoring
```

An up and running postgresdb. Copy the following command for creating a postgres docker container:

```bash
docker run -p 8098:5432  --name monitoring-backend-db --network monitoring -e POSTGRES_PASSWORD=password -e POSTGRES_USER=monitoring -e POSTGRES_DB=monitoring -d postgres
```

An up and running kafka. Therefore the [docker-compose-kafka.yml](../../docker/docker-compose-kafka.yml) can be used.

It is of course possible to run the database and Kafka directly on the host system. Make sure that the monitoring service has the necessary authorization to access the database. Check that the network settings are correct. Changes can be made in the .env file.

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
docker buildx build -t monitoring-backend:latest .
```

**THIS WILL GENERATED SECURITY ISSUES IN YOUR NETWORK**
However, it is necessary to gain access to Kafka if it is not running on the same Docker network.
Run the container:
```bash
docker run --name monitoring-backend --network host -p 8084:8084 -d monitoring-backend 
```

If the broker and the database are on the same network, the container can be started with 
the following command:
```bash
docker run --name monitoring-backend --network monitoring -p 8084:8084 -d monitoring-backend 
```