# How to run in a Virtual Machine
It is assumed that the target is a debian/ ubuntu based VM and the package manager `apt` is available.

**IMPORTANT:**  
For all configuration please refer to the README information in the root directory `./README.md`.

## Install build tools
When running the services with docker skip the first steps and jump to the last step which describes how to install docker. Otherwise skip the docker section and follow the other instructions.

### Git 
```
apt install git
```

#### Node
Follow the instructions of the official nodesource repository: [https://github.com/nodesource/distributions](https://github.com/nodesource/distributions)
1. Download and import GPG key
```bash
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | sudo gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
```

2. Create deb repository
```bash
NODE_MAJOR=20
echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | sudo tee /etc/apt/sources.list.d/nodesource.list
```

3. Run Update and install
```bash
sudo apt-get update
sudo apt-get install nodejs -y
```

#### Golang
Install via apt
```bash
sudo apt install golang 
```

#### Docker (optional)
Required for a fast and easy docker setup when no native build of the services and installation of the db is desired.
In case something is not working follow the instruction on the official docker website [https://docs.docker.com/engine/install/ubuntu/](https://docs.docker.com/engine/install/ubuntu/#prerequisites)


Run the following commands to install docker on the host:
```bash
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```

## Database
Make sure to install postgresdb via docker or native.
For the docker installation please refer to the `README` provided int the backend directory `./backend/README.md`.

Otherwise install via apt:
```bash
sudo sh -c 'echo "deb https://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get -y install postgresql

```

Example default configuration for the database:
```bash
POSTGRES_USER="usermanagement"
POSTGRES_PASSWORD="password"
POSTGRES_DB="usermanagement"
SSL_TLS="disable"
TIMEZONE="Europe/Berlin"
```

When running with a native installation please make sure the configuration in `backend/.env`is the same as the database.

## Kafka
### Native install
Try follow the official instructions, keep in mind to install zookeeper as well.

### Docker
An example docker-compose can be found [here](https://github.com/D4ykoo/travelplatform-case-m2/blob/develop/docker/docker-compose-kafka.yml), just run it and configure as needed:
```bash
docker compose up -d
```

## Run the services by building on the host
#### 1. Clone the repository
```bash
https://github.com/D4ykoo/travelplatform-case-m2.git
```
#### 2. Install and run backend
```bash
cd ./travelplatform-case-m2/usermanagement/backend
# install the packages
go install
# create the binary
go build 
# run the binary 
./usermanagement
```

#### 3. Install and run frontend
For the quick serve, make sure the `backend/.env` configuration has the same port in the `DOMAIN` entry:
```bash
cd ./travelplatform-case-m2/usermanagement/frontend
npm run dev
```

```bash
cd ./travelplatform-case-m2/usermanagement/frontend
npm run build
```
Now the whole application is located in the dist/ directory.<br>
The application can be served by any desired webserver by coping the whole directory and renaming it to e.g. usermanagement. 

##### Alternatively
```bash
cd ./travelplatform-case-m2/usermanagement/frontend
npm run dev
```

## Run the usermanagement services by building on host
#### 1. Clone the repository if not already done
```bash
git clone https://github.com/D4ykoo/travelplatform-case-m2.git
cd ./travelplatform-case-m2/usermanagement/

```

#### 2. Run the docker compose
Optional configure the env vars in the [docker-compose.yml](docker-compose.yml)
```bash
docker compose up -d
```
