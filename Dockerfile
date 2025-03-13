FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download & go mod verify

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
