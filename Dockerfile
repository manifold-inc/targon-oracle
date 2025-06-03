FROM golang:alpine AS build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN GOOS=linux go build -o server ./cmd/api 

FROM alpine:3.9 
WORKDIR /app
RUN apk add ca-certificates
COPY --from=build /app/server server
CMD ["/app/server"]
