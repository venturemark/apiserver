FROM alpine:3.13.3

RUN apk add --no-cache ca-certificates

ADD ./apiserver /apiserver

ENTRYPOINT ["/apiserver"]
