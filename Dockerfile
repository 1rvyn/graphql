# Start from the latest golang base image
FROM golang:1.19

# Add Maintainer Info
LABEL maintainer="Irvyn Hall irvynhall@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

ENV SECRET_KEY=$SECRET_KEY \
    HASH_KEY=$HASH_KEY \
    DB_SERVER=$DB_SERVER \
    DB_USER=$DB_USER \ 
    DB_PASSWORD=$DB_PASSWORD \
    DB_NAME=$DB_NAME \ 
    PORT=$PORT 

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./main"]
