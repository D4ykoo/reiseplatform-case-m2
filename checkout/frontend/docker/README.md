# Run the checkout backend
### Configuration
In the environment section of the corresponding docker-compose<-xx>.yml files.  

### 1. Prerequisites
When running with backend functionallity make sure the backend is running in some way, e. g. using docker.

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