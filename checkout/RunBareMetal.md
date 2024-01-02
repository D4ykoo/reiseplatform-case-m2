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
cd frontend
npm install
ng serve
```

#### Production

```bash
cd frontend
npm install
# either
ng build
# or
npm build
```

Now the whole application is located in the dist/frontend/browser directory.  
The application can be served by any desired webserver after coping the whole directory to e. g. `var/www/html/`.
