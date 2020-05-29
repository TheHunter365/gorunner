FROM golang:alpine AS builder

WORKDIR /usr/src/
COPY . .
RUN go build .

FROM alpine

WORKDIR /usr/app
COPY --from=builder /usr/src/gorunner gorunner
EXPOSE 8080
ENTRYPOINT ["./gorunner"]
