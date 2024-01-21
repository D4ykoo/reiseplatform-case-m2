# How to run 
### Configuration
There are three different ways for configuration:

#### Development
Change the values of [`src/environment/environment.ts`](src/environment/enviroment.ts)

#### Production
The following configuration methods make it possible to set the environment variables in the Dockerfile. That is the reason why this was implemented for production.  

The first option is via the [`src/assets/env.js`](src/assets/env.js) file.  
Or by exporting the following and replacing [`env.js`](src/assets/env.js) with the template:
```bash
# example export API_URL
export API_URL="https://localost:8084/api/v1";

# Replace variables in env.js
envsubst < assets/env.template.js > assets/env.js
```

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
Make sure node_modules exists, if not run `npm install`

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

Optional pass the env file path as argument when chaning the default path:
```bash
docker run --env ENV_FILE="./example.docker.env" --name checkout-frontend --network checkout -p 8083:80 -d checkout-frontend
```