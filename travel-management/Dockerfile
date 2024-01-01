### Frontend build ###
FROM node:21.5.0 AS build
WORKDIR /usr/src/app
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install
COPY frontend .
RUN npm run build
### Backend build ###
FROM golang:alpine3.19 
WORKDIR /app
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download
COPY ./backend ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /travelmanagement
### RUN ###
RUN mkdir html
COPY --from=build /usr/src/app/dist/frontend/browser ./html
RUN apk --update add envsubst
CMD ["/bin/sh",  "-c",  "envsubst < ./html/assets/env.template.js > ./html/assets/env.js && /travelmanagement"]