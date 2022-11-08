FROM golang:1.17.7-buster AS builder
WORKDIR /proxy
ADD vendor ./vendor
ADD cmd ./cmd
ADD go.mod .
ADD go.sum .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix nocgo -o /app cmd/main.go

FROM scratch
COPY --from=builder /app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /tmp /tmp

EXPOSE 8080
WORKDIR /
ENTRYPOINT ["/app"]
