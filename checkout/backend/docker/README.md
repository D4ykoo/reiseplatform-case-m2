# Run the checkout backend
### Configuration
In the environment section of the corresponding compose files.  
The kafka topic creation can be set in the kafka compose file e. g. `--topic=checkout` in the command section of the generator service. 

### 1. Kafka if not already set up
```bash
docker compose -f docker-compose-kafka.yml up -d 
```

### 2. Registry and local image
Decide if local building is desired or not. 

#### 2.1. When using the registry image:
```bash
docker compose up -d
```

#### 2.2. If building the image locally:
```bash
docker compose -f docker-compose-build.yml up -d
```
then run the services with:

```bash
docker compose -f docker-compose-local.yml up -d
```