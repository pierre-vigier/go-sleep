FROM golang:1.14 AS builder

RUN useradd -u 10001 sleeper

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64

WORKDIR /src
COPY ./ ./

RUN go build -o /go-sleep .

FROM scratch AS final

COPY --from=builder /go-sleep /
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER sleeper

ENTRYPOINT ["/go-sleep"]
