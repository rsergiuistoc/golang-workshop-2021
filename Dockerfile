FROM golang:alpine

# Define workdir
WORKDIR /todo-service

# Copy mod and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# COpy the code into the container
COPY . .

# Build the application
RUN go build -v -o todo-service cmd/main.go

# Expose pod
EXPOSE 5000

# Command to run
ENTRYPOINT ["./todo-service"]