FROM golang:1.13.1-alpine

ENV GO111MODULE=on

RUN apk update && \
    apk add emacs curl git && \
    go get github.com/pilu/fresh

# go library
RUN  go get github.com/go-sql-driver/mysql \
        && go get github.com/labstack/echo

CMD ["go", "run", "main.go"]