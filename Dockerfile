FROM golang

ADD . /go/src/server

ENV APP_HOME /go/src/server

WORKDIR $APP_HOME

RUN go build -o server cmd/main.go

EXPOSE 8080 

ENTRYPOINT ./server