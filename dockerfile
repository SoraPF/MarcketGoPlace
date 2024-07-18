FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bin/Marketplace

EXPOSE 3000

CMD ["/app/bin/Marketplace"]
