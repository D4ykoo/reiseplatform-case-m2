# How to run 
### Configuration
There are three different ways for configuration:

#### Development
Change the values of [`src/environments/environment.ts`](src/environments/enviroment.ts)

#### Production
The following configuration methods make it possible to set the environment variables in the Dockerfile. That is the reason why this was implemented for production.  

The first option is via the [`src/assets/env.js`](src/assets/env.js) file.  
Or by exporting the following and replacing [`env.js`](src/assets/env.js) with the template:
```bash
# example export API_URL
export Monitor_API="https://localost:8088/api/v1";

# Replace variables in env.js
envsubst < assets/env.template.js > assets/env.js
```
The following variables can be set:
- Monitor_API (Backend API)
- Checkout_URL
- Login_URL
- TRAVEL_URL

## Bare Metal
```bash
npm install 
# now run 
ng serve --port 8087 
# or for prod with a webserver like nginx
npm run build
```
Now the whole application is located in the dist/ directory.<br>
The application can be served by any desired webserver after coping the whole directory to e. g. `var/www/html/`.<br>
The backend sets the correct HTTP header when the webclient requests arrive via localhost:8087. This prevents CORS errors

## Docker
Create the network if it does not already exist:

Build the image:
```bash
docker buildx build -t monitoring-frontend:latest .
```

Run the container:
```bash
docker run --name monitoring-frontend -p 8087:8087 -d monitoring-frontend
```

Optional pass the env file path as argument when chaning the default path:
```bash
docker run --env ENV_FILE="./example.docker.env" --name checkout-frontend --network checkout -p 8087:8087 -d checkout-frontend
