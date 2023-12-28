FROM oven/bun:1 as base
WORKDIR /usr/src/app

# cache for faster future builds
FROM base AS install
RUN mkdir -p /tmp/dev
COPY package.json bun.lockb /tmp/dev/
RUN cd /tmp/dev && bun install --frozen-lockfile

# without dev dependencies
RUN mkdir -p /tmp/prod
COPY package.json bun.lockb /tmp/prod/
RUN cd /tmp/prod && bun install --frozen-lockfile --production

# copy node_modules from tmp directory
# then all non ignored files
FROM base AS prerelease
COPY --from=install /tmp/dev/node_modules node_modules
COPY . .

# build
ENV NODE_ENV=production
RUN bun run build
COPY . .

# for debugging purposes 
RUN dir -s    
RUN dir build   

# copy production dependencies and source code into final image
FROM base AS release
RUN file="$(ls -1 /usr/src/app/)" && echo $file
COPY --from=install /tmp/prod/node_modules node_modules
COPY --from=prerelease /usr/src/app/build .

# run the app
USER bun
EXPOSE 3000/tcp
ENTRYPOINT [ "bun", "run", "index.js"]