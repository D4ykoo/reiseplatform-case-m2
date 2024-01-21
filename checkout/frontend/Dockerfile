FROM node:latest as builder
WORKDIR /usr/src/app
ARG frontend_path=./

ARG env_file=./example.docker.env
ENV ENV_FILE=$env_file

COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build
### Run ###
FROM nginx:stable-alpine3.17
COPY --from=builder /usr/src/app/dist/frontend/browser /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 8083
CMD ["/bin/sh",  "-c",  "envsubst < /usr/share/nginx/html/assets/env.template.js > /usr/share/nginx/html/assets/env.js && exec nginx -g 'daemon off;'"]
