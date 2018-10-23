
FROM golang:alpine as builder
RUN apk --update add ca-certificates
ARG srcdir=/go/src/github.com/citizen428/unsavory/
RUN mkdir -p $srcdir
ADD . $srcdir
WORKDIR $srcdir
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags '-extldflags "-static"' -o /unsavory ./cmd/unsavory

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt \
     /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /unsavory /
ENTRYPOINT ["/unsavory"]
