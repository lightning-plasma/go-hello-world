FROM golang:1.14.1-alpine3.11

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main main.go

ENTRYPOINT ["/app/main"]
EXPOSE 80
