FROM rust:latest as build

### Build ###
WORKDIR /app

COPY . .

# build for release
RUN cargo build --release
RUN rm -f target/release/deps/monitoring*

# our final base
FROM debian:bookworm-slim
WORKDIR /app
RUN apt-get -y update && apt-get install -y --no-install-recommends libpq-dev

COPY --from=build /app/target/release/monitoring-service /app/monitoring-service/.env ./
CMD ["./monitoring-service"]