# Run the checkout backend
All-in-One checkout docker compose solution. Contains:
* Kafka
* DB
* Backend
* Frontend

Optional: There is a docker compose without kafka. Make sure the kafka instance is accesible by the containers when running this file with:
```bash
docker compose -f docker-compose-without-kafka.yml up -d 
```
An example kafka docker compose can be found in the backend directory.

### 1. Configuration
In the environment section of the corresponding docker-compose<-xx>.yml files.  
The kafka topic creation can be set in the kafka compose file e. g. `--topic=checkout` in the command section of the generator service. 

### 2. Registry and local image
Decide if local building is desired or not. 

#### 2.1. When using the registry image:
```bash
docker compose up -d
```

#### 2.2. If building the image locally beforehand:
```bash
docker compose -f docker-compose-local.yml up -d
```