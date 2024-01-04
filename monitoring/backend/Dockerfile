### Build ###
FROM rust:latest as build

WORKDIR /app

COPY . .

RUN cargo build --release
RUN rm -f target/release/deps/monitoring*

### RUN ###
FROM debian:bookworm-slim
WORKDIR /app

RUN apt-get -y update && apt-get install -y --no-install-recommends libpq-dev
COPY --from=build /app/target/release/monitoring-service /app/monitoring-service/.env ./

CMD ["./monitoring-service"]