# How to run 
### Configuration
In `src/assets/config.ts`

## Bare Metal
```bash
npm run dev 
# or for prod
```

## Docker
Create the network if it does not already exist:
```bash
   docker network create usermanagement
```

Build the image:
```bash
docker buildx build -t usermanagement-frontend:latest .
```

Run the container:
```bash
docker run --name usermgm-frontend --network usermanagement -p 8081:80 -d usermanagement-frontend
```

```bash
```

```bash
```