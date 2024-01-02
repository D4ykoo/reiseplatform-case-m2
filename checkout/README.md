# Checkout/ Payment Microservice

## Prerequisites
* A postgresql database where the configuration and credentials are known

## How to run - bare metal

### Configuration
All the configuration for the backend is done in the .env of the file. See comments for some additional information.

### Backend
#### Development
```bash
cd backend
cargo run 
# alternatively use hot reload with
cargo install cargo-watch
cargo watch -x run
```

#### Production
```bash
cd backend
cargo build --release
./target/release/checkout-backend
```

### Frontend
#### Development
```bash
bun --bun run dev
```

#### Production
First change the directory and add the recommended SvelteKit adapter:
```bash
cd frontend
bun add -D svelte-adapter-bun
```
Then import it by replacing the existing one in `svelte.config.js` 

```bash
 - import adapter from "@sveltejs/adapter-auto";
 + import adapter from "svelte-adapter-bun";
```

Finally build the frontend:
```bash
bun run build
```

## How to run - Dockerfile
Note: due to the microservice architecture the frontend __and__ backend have a `Dockerfile`. The `Dockerfile` at the root of this directory is the All-in-One Solution, only the frontend will be exposed.

Since the `docker build` is marked as deprecated the new `cli` extension is used:
```bash
cd frontend
docker buildx build --pull -t checkout-frontend .
```
Now run the image:
```bash
docker run -p 3000:3000 checkout-frontend
```
optional append the arg for the frontend paths when having different configuration on the system and not the default structure like after cloning the repository:
```bash
# trailing slash for the path is needed
docker run -p 3000:3000 checkout-frontend --build-arg="/path/to/dir/"
```
#### Backend
Build image:
```bash
cd backend
docker buildx build -t checkout-backend .
```

Run:
```bash
docker run -p 8071:8071 checkout-backend
```


## Technology Stack
* Svelte + SvelteKit + bun + TailwindCSS + daisyUI
* Rust + Actix Web + Diesel

## Additional Information
* Project management done with self-hosted [Leantime](https://github.com/Leantime/leantime)
* Verification of the JWT: 
  * Username included 
  * -> extract JWT via browser dev tools and paste into [https://jwt.io/](https://jwt.io/)
  * Different users have different usernames