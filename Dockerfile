FROM golang:1.23.1-alpine AS builder

# Set working directory inside the container
WORKDIR /app
# Copy code from current directory into container
COPY . .

# Build the Go app into an executable named "main"
RUN go get && \
    go build -a -o main .

FROM alpine:3.10
COPY --from=builder /app/main .

# Run the built executable
ENTRYPOINT ["./main"]
