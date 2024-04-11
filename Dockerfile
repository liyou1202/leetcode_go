FROM golang:1.22-alpine AS builder

COPY . /src

WORKDIR /src

RUN go build main.go


FROM alpine:3.17.0

COPY --from=builder /src/main /src/main

ENTRYPOINT [ "/src/main"]
