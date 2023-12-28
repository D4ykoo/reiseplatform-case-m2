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

## Technology Stack
* Svelte + SvelteKit + bun + TailwindCSS + daisyUI
* Rust + Actix Web + Diesel

## Additional Information
* Project management done with self-hosted [Leantime](https://github.com/Leantime/leantime)
* Verification of the JWT: 
  * Username included 
  * -> extract JWT via browser dev tools and paste into [https://jwt.io/](https://jwt.io/)
  * Different users have different usernames