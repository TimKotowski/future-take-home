FROM golang:1.23.2-bullseye as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/future ./cmd

FROM debian:12

COPY --from=builder /build/bin/future /

EXPOSE 8080

COPY migrations /migrations

CMD [ "/future", "future" ]