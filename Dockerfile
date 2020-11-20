FROM golang:1.14 as build-env

RUN mkdir -p $GOPATH/src/nano-world-backend-golang
COPY . $GOPATH/src/nano-world-backend-golang/
WORKDIR $GOPATH/src/nano-world-backend-golang/

RUN go get ./...


RUN go build -o main
CMD ["./main"]