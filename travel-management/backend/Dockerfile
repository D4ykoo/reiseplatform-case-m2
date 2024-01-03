FROM golang:alpine as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /travelmanagement

FROM alpine:latest
WORKDIR /app
COPY --from=build travelmanagement .
COPY .env . 

EXPOSE 8082
CMD ["./travelmanagement"]