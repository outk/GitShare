FROM golang:1.18.2

WORKDIR /app

COPY go.mod go.sum ./
COPY app/ ./

RUN go mod download
RUN go build -o /gifshare

EXPOSE 8080

CMD [ "/gifshare" ]