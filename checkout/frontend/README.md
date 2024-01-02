# How to run 
### Configuration
In [src/environment/environment.prod.ts](src/environment/enviroment.prod.ts), when running in dev mode change the config on the non prod file.

## Bare Metal
```bash
npm install 
# now run
npm run dev 
# or for prod with a webserver like nginx
npm run build
```
Now the whole application is located in the dist/ directory.<br>
The application can be served by any desired webserver after coping the whole directory to e. g. `var/www/html/`.

## Docker
Create the network if it does not already exist:
```bash
docker network create checkout
```

Build the image:
```bash
docker buildx build -t checkout-frontend:latest .
```

Run the container:
```bash
docker run --name checkout-frontend --network checkout -p 8083:80 -d checkout-frontend
```