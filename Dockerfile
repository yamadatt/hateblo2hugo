#FROM golang:1.17.7-alpine3.15
#FROM golang:latest
FROM golang:1.21.1-alpine3.17


WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./

#RUN go mod download

COPY ./ ./

RUN go mod download

RUN go build -o /main


#CMD ["/main"]
ENTRYPOINT ["/main"]
