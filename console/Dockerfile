# Path: /console/Dockerfile
# Purpose: Defines the build process for the Vite/React frontend console service for local development.
# Note: This file is part of Node 4's implementation.

# Use the official Node.js image.
FROM node:20.15.0-alpine

# Update base packages to patch potential vulnerabilities.
RUN apk upgrade --no-cache

# Set the working directory inside the container.
WORKDIR /app

# Copy package.json and lock file to leverage Docker's dependency caching.
COPY package*.json ./

# Install dependencies.
RUN npm install

# Copy the rest of the application source code.
# The `volumes` in docker-compose.yml will overwrite this for live-reloading.
COPY . .

# The command to start the Vite development server.
# The `--host` flag is crucial to expose the server outside the container's network interface.
CMD ["npm", "run", "dev", "--", "--host"]