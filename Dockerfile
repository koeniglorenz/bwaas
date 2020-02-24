FROM golang:1.11-alpine

EXPOSE 8080

RUN mkdir -p $GOPATH/src/github.com/koeniglorenz/bwaas
WORKDIR $GOPATH/src/github.com/koeniglorenz/bwaas

COPY . .

RUN go build -o bwaas ./cmd/main.go

CMD ["./bwaas", "--buzzwords", "./buzzwords.json"]
