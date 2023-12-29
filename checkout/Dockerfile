FROM oven/bun:1 as base
WORKDIR /usr/src/app
ARG frontend_path=frontend/

# cache for faster future builds
FROM base AS install
RUN mkdir -p /tmp/dev
COPY ${frontend_path}package.json ${frontend_path}bun.lockb /tmp/dev/
RUN cd /tmp/dev && bun install --frozen-lockfile

# NOTE: this step is quite nice BUT
# not suitable for this use case with svelte kit
# otherwise necessary packages are missing (strange behaviour)
# without dev dependencies
RUN mkdir -p /tmp/prod
COPY ${frontend_path}package.json ${frontend_path}bun.lockb /tmp/prod/
RUN cd /tmp/prod && bun install --frozen-lockfile --production

# copy node_modules from tmp directory
# then all non ignored files
FROM base AS prerelease
COPY --from=install /tmp/dev/node_modules node_modules
COPY . .

# build
ENV NODE_ENV=production
RUN cd ${frontend_path} && bun run build
COPY ${frontend_path}build ./build/

# copy production dependencies and source code into final image
FROM base AS release
COPY --from=install /tmp/dev/node_modules node_modules
COPY --from=prerelease /usr/src/app/build .

# run the app
USER bun
EXPOSE 8085/tcp
ENTRYPOINT [ "bun", "run", "index.js"]