FROM golang:1.15

COPY . /go/src/app

WORKDIR /go/src/app/cmd/server

RUN go build -o server main.go

CMD ["./server"]
