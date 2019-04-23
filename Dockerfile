# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.11

# Add Maintainer Info
LABEL maintainer="Valeri Zimin"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/go-packg/dev/cmd/cloudnativego/go-primer/npcs

# Copy the local package files to the container's workspace.
# ADD . /go/src/github.com/go-packg/dev/cmd/cloudnativego/go-primer/npcs

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Build the npcs command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
# RUN go install github.com/go-packg/dev/cmd/cloudnativego/go-primer/npcs

# Install the package
RUN go install -v ./...

# Document that the service listens on port 8080.
# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the npcs command by default when the container starts.
# ENTRYPOINT /go/bin/npcs

# Run the executable
CMD ["npcs"]