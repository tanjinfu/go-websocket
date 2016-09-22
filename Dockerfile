FROM golang:1.6
EXPOSE 8080

MAINTAINER Tan Jinfu

ADD main.go /var/websocket/main.go
ADD index.html /var/websocket/index.html
ADD jquery-2.1.4.min.js /var/websocket/jquery-2.1.4.min.js

RUN mkdir -p /go/src/golang.org/
ADD golang.org /go/src/golang.org/

WORKDIR "/var/websocket"
CMD ["go", "run", "/var/websocket/main.go"]