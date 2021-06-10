FROM golang:1.16

# Set current working directory
WORKDIR /app

# Now, copy the source code
COPY . .

# Download all dependencies.
RUN go build -o main main.go
# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

# Finally our multi-stage to build a small image
# Start a new stage from scratch

EXPOSE 8003

# Run executable
CMD ["./main"]