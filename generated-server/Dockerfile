# Stage 1: Build the Go application
FROM golang:1.22 AS build
WORKDIR /app

# Copy Go modules manifests
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY ./go ./go
COPY ./internal ./internal
COPY ./certs ./certs
COPY ./main.go ./

# Set CGO_ENABLED to 1 for building with C dependencies
ENV CGO_ENABLED=1

# Run Go build
RUN go mod tidy
RUN go build -o serv .

# Stage 2: Create the minimal runtime image
FROM alpine:latest AS runtime

# Install necessary libraries for CGO
RUN apk add --no-cache libc6-compat

# Copy the Go binary and certificates from the build stage
COPY --from=build /app/serv /serv
COPY --from=build /app/certs /certs

# Expose the port the application will run on
EXPOSE 3333/tcp

# Run the Go binary
ENTRYPOINT ["/serv"]
