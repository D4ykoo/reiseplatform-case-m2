

# How to run 
### Configuration
Can be applied in the .env file.

### Prerequisites
Create the network if it does not already exist:
```bash
   docker network create usermanagement
```

An up and running postgresdb. Copy the following command for creating a postgres docker container:

```bash
docker run -p 8092:5432  --name usermgm-backend-db --network usermanagement -e POSTGRES_PASSWORD=password -e POSTGRES_USER=usermanagement -e POSTGRES_DB=usermanagement -d postgres
```
## Bare Metal
```bash
go run ./main.go
```

## Docker
Build the image:
```bash
docker buildx build -t usermanagement-backend:latest .
```

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