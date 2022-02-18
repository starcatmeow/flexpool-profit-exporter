FROM golang:1.16-alpine AS builder
RUN apk update && apk add --no-cache git upx ca-certificates
WORKDIR /src/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o exporter
RUN upx --lzma exporter
RUN chown nobody:nogroup exporter

FROM scratch
WORKDIR /app
USER nobody
COPY --from=builder /etc/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/exporter /app/exporter
EXPOSE 9091
CMD ["/app/exporter"]
