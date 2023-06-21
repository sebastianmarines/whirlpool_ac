FROM golang:1.20.5

ADD . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o app


ENTRYPOINT /app/app

# Build final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /app/app .

CMD ["./app"]