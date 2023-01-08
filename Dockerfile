FROM golang:1.14 AS builder
WORKDIR /go/src/githum.com/mig-elgt/osrm-fetcher
ADD . .
RUN go test --cover -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o /bin/app main.go

FROM alpine:3.4
MAINTAINER Miguel Angel Galicia
RUN apk --no-cache --update add ca-certificates

COPY --from=builder /bin/app /usr/local/bin/app
RUN chmod +x /usr/local/bin/app
CMD ["app"]
