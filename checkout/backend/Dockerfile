# The rust alpine image can cause some erros so the official debian image is used
FROM rust:latest as builder
WORKDIR /app
ARG backend_path=./

# FROM builder AS install
COPY . .

FROM builder AS build
RUN cargo build --release

FROM builder AS release
COPY --from=build /app/target/release/checkout-backend /usr/bin/server

EXPOSE 8071/tcp
CMD ["server"]