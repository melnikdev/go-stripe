FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download & go mod verify

COPY . .

RUN go build -o main cmd/api/main.go

EXPOSE 8080

CMD ["./main"]
