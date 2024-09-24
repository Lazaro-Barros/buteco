FROM golang:latest as build

WORKDIR /app

# Copy the Go module files
COPY go.mod . 
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

# Build the application as a static binary
ENV CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:latest as run
WORKDIR /app

# Install certificates for HTTPS requests
RUN apk add --no-cache ca-certificates

# Copy the application executable from the build image
COPY --from=build /app/app /app/app

# Ensure the executable has the right permissions
RUN chmod +x /app/app

EXPOSE 8080
CMD ["./app"]
