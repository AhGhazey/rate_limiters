FROM golang:1.21-alpine3.18 as builder
LABEL authors="ahmedghazey"

USER root
RUN apk update && apk add g++ make openssh-client bash git

WORKDIR /go/src/github/ahghazey/rate_limiter

ADD Makefile .
ADD go.mod .
ADD go.sum .

COPY . .
RUN make install
RUN make mod
COPY . .
RUN make build

# Main image
FROM golang:1.21-alpine3.18
WORKDIR app
RUN mkdir -p /etc/config/rate_limiter

COPY --from=builder /go/src/github/ahghazey/rate_limiter/rate_limiters .
COPY --from=builder /go/src/github/ahghazey/rate_limiter/assets /etc/config/rate_limiter

ENTRYPOINT ["/bin/sh", "-c", "/go/app/rate_limiters"]