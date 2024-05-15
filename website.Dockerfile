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


FROM node:22-alpine

WORKDIR /app

COPY --from=frontend-build /src .

CMD ["npm", "run", "dev"]
