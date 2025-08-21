FROM golang:1.25-bookworm AS build-env

# Create and change to the app directory.
WORKDIR /app

#disable crosscompiling
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

#build the binary with debug information removed
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o /azure-volume-populator github.com/pdok/azure-volume-populator

FROM scratch
EXPOSE 8080

WORKDIR /
ENV PATH=/

COPY --from=build-env /azure-volume-populator /

ENTRYPOINT ["/azure-volume-populator"]