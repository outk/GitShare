FROM golang:1.18-alpine

ENV goroot=/usr/local/go/src

WORKDIR ${goroot}/gifshare/app

COPY go.mod go.sum ./
RUN go mod download

COPY app/ ./
RUN go build -o /gifshare

EXPOSE 8080

CMD [ "/gifshare" ]