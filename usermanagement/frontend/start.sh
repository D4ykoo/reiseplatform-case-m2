#!/bin/bash
set -e

chmod -R 777 /app
cd /app
./import-meta-env -x .env

cd /app/dist/
cp -r * /usr/share/nginx/html/
nginx -g "daemon off;"