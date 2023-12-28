FROM node:21.5.0 AS build
WORKDIR /usr/src/app
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install
COPY frontend .
RUN npm run build
### STAGE 2: Run ###
FROM nginx:alpine
COPY --from=build /usr/src/app/dist/frontend/browser /usr/share/nginx/html