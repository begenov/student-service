FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o student-service ./cmd/main.go

CMD ["./student-service"]
