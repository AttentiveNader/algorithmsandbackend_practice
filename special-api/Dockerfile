FROM golang:1.8

ENV GOPATH /go
ENV GOBIN /go/bin

WORKDIR /home/nader/go/src/backend_practice/special-api
COPY . .

RUN go get -d -v github.com/satori/go.uuid
RUN go get -d -v github.com/gorilla/websocket
RUN go install /home/nader/go/src/backend_practice/special-api/main.go

EXPOSE 8080


CMD ["go","run","main.go"]