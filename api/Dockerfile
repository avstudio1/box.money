# Path: /api/Dockerfile
# Purpose: Defines the build process for the Go backend API service.
# Note: This version is optimized for local development with live-reloading using 'air'.

# Use the official Go image.
FROM golang:1.21-alpine

# Set the working directory inside the container.
WORKDIR /app

# Install 'air' for live-reloading by downloading the pre-compiled binary.
# Also upgrade existing packages to patch security vulnerabilities.
RUN apk upgrade --no-cache && apk add --no-cache curl
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /go/bin

# Copy go.mod and go.sum files to leverage Docker's dependency caching.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code.
# The volume mount in docker-compose.yml will keep this in sync.
COPY . .

# The command to start the application using air.
# Air will watch for file changes and automatically recompile and restart the server.
CMD ["air"]