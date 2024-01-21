# How to run in a Virtual Machine
It is assumed that the target is a debian/ ubuntu based VM and the package manager `apt` is available.

**Important:** When not using docker (compose) for the services. Make sure to configure the `.env` files to fit your needs. 

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

4. Install Angular
```bash
sudo npm install -g @angular/cli
```
#### Golang
Install build dependencies:
```bash
sudo apt install build-essential libpq -y
```
Install Golang via apt:
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
POSTGRES_USER="travelmanagement"
POSTGRES_PASSWORD="password"
POSTGRES_DB="travelmanagement"
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
cd ./travelplatform-case-m2/travel-management/backend
# install the packages
go install
# create the binary
go build 
# run the binary 
./travelmanagement
```

#### 3. Install and run frontend
For quick serve:
```bash
cd ./travelplatform-case-m2/travel-management/frontend
ng serve --port 8085
```
It is important to use port number 8085. Otherwise your API calls will be blocked by CORS.

For building:
```bash
cd ./travelplatform-case-m2/travel-management/frontend
ng build
```
Now the whole application is located in the dist/ directory.<br>
The application can be served by any desired webserver by coping the whole directory and renaming it to e.g. usermanagement. 


## Run the travel management services by building on host
#### 1. Clone the repository if not already done
```bash
git clone https://github.com/D4ykoo/travelplatform-case-m2.git
cd ./travelplatform-case-m2/travel-management/
```

#### 2. Run the docker compose
Optional configure the env vars in the [/docker/docker-compose-travelmanagement.yml](../../docker/docker-compose-travelmanagement.yml)
```bash
docker compose up -d
```
