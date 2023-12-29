FROM golang:alpine as builder
WORKDIR /app

ARG env_file=../.env
ENV ENV_FILE=$env_file

COPY go.mod .
RUN go mod download

COPY . .
COPY ${env_file} /usermanagement/.env
RUN go build -o /usermanagement/usermanagement-backend

RUN ls -la /usermanagement

EXPOSE 8081
CMD ["/usermanagement/usermanagement-backend"]