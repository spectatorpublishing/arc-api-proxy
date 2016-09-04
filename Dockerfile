# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/arc-api-proxy
WORKDIR /go/src/arc-api-proxy
# Build the golang-docker command inside the container.
RUN go install arc-api-proxy
ENTRYPOINT /go/bin/arc-api-proxy
# Run the golang-docker command when the container starts.

# http server listens on port 8080.
