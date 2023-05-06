# syntax=docker/dockerfile:1
FROM golang:1.20

# copy destination
WORKDIR /app

# Copy Go modules
COPY app/go.mod .
COPY app/go.sum .
RUN go mod download

# Copy files and folders
COPY app/utils ./utils
COPY app/main.go .

# Build the application
RUN go build .

# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can (optionally) document in the Dockerfile what ports
# the application is going to listen on by default.
EXPOSE 6000

# Run the application
CMD ["./app -p 6000"]
CMD tail -f /dev/null
# docker build --tag hermes .
# docker run -p 6000:6000 hermes