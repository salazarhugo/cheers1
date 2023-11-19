FROM golang:1.19

# Copy the local library into the container
#COPY ../../gen/go ./

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.build.mod ./go.mod
COPY go.sum ./
RUN go mod download


# Copy the rest of the application code
COPY ./ ./

# Build the Go application
RUN go build -o myapp