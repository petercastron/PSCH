# Build psch in a stock Go builder container
FROM golang:1.16-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /psch
RUN cd /psch && make psch

# Pull psch into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /psch/build/bin/psch /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["psch"]
