FROM golang:1.16-alpine

WORKDIR /src/
COPY * /src/
RUN go build -o calc-api

ENTRYPOINT ["/src/calc-api"]