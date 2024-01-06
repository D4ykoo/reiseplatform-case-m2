# How to run
https://github.com/runtime-env/import-meta-env/blob/main/packages/examples/vite-alpine-example/src/main.js

https://medium.com/rockedscience/how-import-meta-env-saved-us-hours-of-deployment-21c7548c3cf4

https://import-meta-env.org/guide/getting-started/runtime-transform.html

https://github.com/runtime-env/import-meta-env/blob/main/packages/examples/docker-starter-example/start.sh

### Configuration

In the [.env](.env) file of the this frontend directory. It will also affect the docker environments.

## Bare Metal

```bash
npm run dev
# or for prod with a webserver like nginx
npm run build
```

Now the whole application is located in the dist/ directory.<br>
The application can be served by any desired webserver after coping the whole directory to e. g. `var/www/html/`.

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
