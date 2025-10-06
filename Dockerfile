FROM golang:1.25.1-alpine

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o book-service ./cmd/main/main.go

EXPOSE 7000

CMD ["./book-service"]
