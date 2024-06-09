###############################
###### BUILD FRONT END  #######
###############################

FROM node:22 AS frontend-build
WORKDIR /src

# Package.json
COPY package.json ./
COPY package-lock.json ./

# Source files
COPY src ./src
COPY static ./static


# Config files
COPY postcss.config.js ./
COPY svelte.config.js ./
COPY tailwind.config.js ./
COPY tsconfig.json ./
COPY vite.config.ts ./
COPY .npmrc ./

# Build
RUN npm install
RUN npm run build


###############################
######### BUILD SQLC ##########
###############################

FROM debian  AS sqlc-build
COPY --from=sqlc/sqlc:latest /workspace/sqlc /usr/bin/sqlc
WORKDIR /src

COPY sqlc.yaml ./
COPY schema.sql ./
COPY query.sql ./

RUN sqlc generate

###############################
####### BUILD GO-SERVER #######
###############################

FROM golang:alpine AS server-builder
WORKDIR /src

COPY server ./server
COPY go.mod .
COPY go.sum .
COPY --from=sqlc-build /src/server/db/ /server/db/

# NOTE: If third-party deps are installed, then git has to be manually installed here.

ENV GOOS=linux
# not 100% sure if this is necessary
RUN go mod download

RUN go build -o ./server ./server/main.go

###############################
#######       run       #######
###############################
FROM alpine:latest
EXPOSE 8080

COPY --from=frontend-build /src/build /build
COPY --from=server-builder /src/server /bin/server

RUN chmod +x /bin/server

CMD ["/bin/server"]