# Dockerfile for Go services in the backend project
# This Dockerfile is designed to be built from the monorepo root.
# Example: docker build -f projects/backend/Dockerfile --build-arg SERVICE_NAME=auth .

# --- Builder Stage ---
FROM golang:1.25.5-alpine AS builder

WORKDIR /build

# Set GOPROXY for faster dependency downloads in China
ENV GOPROXY=https://goproxy.cn,direct

# To leverage Docker's layer caching, we first copy only the files
# that define our dependencies.
COPY go.mod go.sum ./

# Download all dependencies based on the workspace and module files.
# This step is cached as long as the dependency files don't change.
RUN go mod download

# Now, copy the entire source code.
COPY . .

# Declare the service name to be built, passed as a build argument.
ARG SERVICE_NAME

# Build the Go application.
# - CGO_ENABLED=0 produces a static binary.
# - GOOS=linux ensures it's built for the Alpine base image.
# - -a flag forces rebuilding of packages that are already up-to-date.
# - -o specifies the output file path.
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app/${SERVICE_NAME} ./cmd/${SERVICE_NAME}

# --- Runner Stage ---
FROM alpine:latest

WORKDIR /app

# Argument for the service name, needed again in this stage.
ARG SERVICE_NAME

# Copy the compiled binary from the builder stage.
COPY --from=builder /app/${SERVICE_NAME} .

# Copy the service's configuration files.
# The path is relative to the build context (monorepo root).
COPY resources/configs /data/configs

# Define the entrypoint for the container.
# It executes the service binary, passing the path to the config directory.
ENTRYPOINT ["./${SERVICE_NAME}", "-conf", "/data/configs"]
