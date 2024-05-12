FROM golang:1.22 as build

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN  go build -o /app

EXPOSE 8080

CMD ["/app"]