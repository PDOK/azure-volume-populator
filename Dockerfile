FROM golang:1.25-bookworm AS build-env

# Create and change to the app directory.
WORKDIR /app

# Disable crosscompiling
ENV CGO_ENABLED=0

# Compile linux only
ENV GOOS=linux

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download all && \
    go generate -v ./...

# Copy local code to the container image.
COPY . ./

# Build the binary with debug information removed
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o /azure-volume-populator github.com/pdok/azure-volume-populator/cmd/

# Final image
FROM docker.io/debian:bookworm-slim
EXPOSE 8080

RUN set -eux && \
    apt-get update && \
    apt-get install --no-install-recommends -y \
      ca-certificates=* && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /
ENV PATH=/

COPY --from=build-env /azure-volume-populator /

ENTRYPOINT ["/azure-volume-populator"]