
# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Shohag Rana <shohagrana64@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /api

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build binaries from the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

############# New Stage ################
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /api/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

