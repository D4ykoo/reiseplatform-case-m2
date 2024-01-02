FROM builder AS install
RUN mkdir -p /tmp/dev
COPY ${frontend_path}package.json ${frontend_path}package-lock.json /tmp/dev/
RUN cd /tmp/dev && npm install --frozen-lockfile

RUN mkdir -p /tmp/prod
COPY ${frontend_path}package.json ${frontend_path}package-lock.json /tmp/prod/
RUN cd /tmp/prod && npm install --frozen-lockfile --production

FROM builder AS prerelease
COPY --from=install /tmp/dev/node_modules node_modules
COPY . .

ENV NODE_ENV=production
RUN cd ${frontend_path} && npm run build
COPY ${frontend_path}dist .

FROM nginx:alpine 
COPY --from=prerelease /app/dist/ /usr/share/nginx/html

COPY nginx.conf /etc/nginx/nignx.conf

EXPOSE 8083
CMD ["nginx", "-g", "daemon off;"]