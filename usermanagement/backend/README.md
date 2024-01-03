# How to run
### Configuration
Can be applied in the .env file.
Or when using the docker hub registry image provide the env file via the docker run option `docker run --env-file docker.env.example <...>` in your. The example file can be renamed and the values in it changed to fit yor needs.

### Prerequisites
Make sure the network exists if a seperate network is desired.  
An up and running postgresdb. Copy the following command for creating a postgres docker container:

```bash
docker run -p 8092:5432  --name usermgm-backend-db --network usermanagement -e POSTGRES_PASSWORD=password -e POSTGRES_USER=usermanagement -e POSTGRES_DB=usermanagement -d postgres
```
## Bare Metal
```bash
go run ./main.go
# or build and run 
go build && ./usermanagement
```

## Docker
Create the network if it does not already exist:
```bash
docker network create usermanagement
```

Build the image:
```bash
docker buildx build -t usermanagement-backend:latest .
```
Could be pushed to docker hub with:
```bash
docker tag local-image:tagname new-repo:tagname
docker push new-repo:tagname
```
### Run with local built image
**THIS WILL GENERATED SECURITY ISSUES IN YOUR NETWORK**  
Run the container:
```bash
docker run --name usermgm-backend --network host -p 8082:8082 -d usermanagement-backend 
```

**CURRENTLY NOT WORKING**  
Run the container:
```bash
docker run --name usermgm-backend --network usermanagement -p 8082:8082 -d usermanagement-backend 
```

### Run with docker hub image
```bash
sudo docker run -p 8082:8082 -d dak4408/travelplatform-case-m2:usermanagement-backend
```

## HTTP Routes
| **Method** 	| **Route**        	|
|------------	|------------------	|
| GET        	| /v1/api/users    	|
| GET        	| /v1/api/users:id 	|
| POST       	| /v1/api/users    	|
| PUT        	| /v1/api/users:id 	|
| DELETE     	| /v1/api/users:id 	|
| POST       	| /v1/api/login    	|
| POST       	| /v1/api/register 	|
| PUT        	| /v1/api/reset    	|
| GET        	| /v1/api/logout   	|