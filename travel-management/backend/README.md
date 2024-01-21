# How to run
### Configuration
Can be applied in the .env file.
Or when using the docker hub registry image provide the env file via the docker run option `docker run --env-file docker.env.example <...>` in your. The example file can be renamed and the values in it changed to fit yor needs.

### Prerequisites
Make sure the network exists if a seperate network is desired.  
An up and running postgresdb. Copy the following command for creating a postgres docker container:

```bash
docker run -p 8096:5432  --name travelmgnt-backend-db --network travelmgnt -e POSTGRES_PASSWORD=password -e POSTGRES_USER=travelmanagement -e POSTGRES_DB=travelmanagement -d postgres
```
An up and running kafka. Therefore the [docker-compose-kafka.yml](../../docker/docker-compose-kafka.yml) can be used.

It is of course possible to run the database and Kafka directly on the host system. Make sure that the monitoring service has the necessary authorization to access the database. Check that the network settings are correct. Changes can be made in the .env file.
## Bare Metal
```bash
go run ./main.go
# or build and run 
go build && ./travelmanagement
```

## Docker
Create the network if it does not already exist:
```bash
docker network create travelmgnt
```

Build the image:
```bash
docker buildx build -t travelmanagement-backend:latest .
```

### Run with local built image
**THIS WILL GENERATED SECURITY ISSUES IN YOUR NETWORK**<br>
However, it is necessary to gain access to Kafka if it is not running on the same Docker network. Run the container:  
```bash
docker run --name travelmgnt-backend --network host -p 8086:8086 -d travelmanagement-backend 
```

If the broker and the database are on the same network, the container can be started with 
the following command:
Run the container:
```bash
docker run --name travelmgnt-backend --network travelmanagement -p 8086:8086 -d travelmanagement-backend 
```

### Run with docker hub image
```bash
sudo docker run -p 8086:8086 -d mig3177/travma-travelmanagement-backend
```

## HTTP Routes
| **Method** 	| **Route**        	              |
|------------	|-------------------------------  |
| GET        	| /api/v1/loginstatus             |
| GET        	| /api/v1/hotels                  |
| GET        	| /api/v1/hotels/:id              |
| POST       	| /api/v1/hotels                  |
| PUT        	| /api/v1/hotels/:id              |
| DELETE     	| /api/v1/hotels/:id              |
| POST       	| /api/v1/hotels/:id/travels      |
| GET       	| /api/v1/hotels/:id/travels/:tid |
| PUT        	| /api/v1/hotels/:id/travels/:tid |
| DELETE        | /api/v1/hotels/:id/travels/:tid |
| POST       	| /api/v1/tags                    |
| GET       	| /api/v1/tags                    |
| GET       	| /api/v1/tags/:id                |
| PUT        	| /api/v1/tags/:id                |
| DELETE        | /api/v1/tags/:id                |