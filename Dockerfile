FROM golang:1.21.1-alpine3.17

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o /main


#CMD ["/main"]
ENTRYPOINT ["/main"]
