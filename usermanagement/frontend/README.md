# How to run

### Configuration

In the [.env](.env) file of the this frontend directory. It will also affect the docker environments.

**IMPORTANT** When running *not* in AIO mode (docker compose in the repository root), change the base in [vite.config.ts](vite.config.ts) to something like `/`. The base property represents the `base href="/users/` tag in html.

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
