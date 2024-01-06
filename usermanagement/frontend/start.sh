#!/bin/bash
set -e

chmod -R 777 /app
cd /app
./import-meta-env -x .env

cd /app/dist/
nginx -g "daemon off;"