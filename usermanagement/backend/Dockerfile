FROM golang:alpine as builder
WORKDIR /app

ARG env_file=../.env
ENV ENV_FILE=$env_file

COPY go.mod .
RUN go mod download

COPY . .
COPY ${env_file} /usermanagement/.env
RUN go build -o /usermanagement/usermanagement-backend

FROM alpine:latest
WORKDIR /app
COPY --from=builder /usermanagement/usermanagement-backend .
COPY --from=builder /usermanagement/.env . 

EXPOSE 8083
CMD ["./usermanagement-backend"]