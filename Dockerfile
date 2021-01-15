FROM alpine:3.13.0

RUN apk add --no-cache ca-certificates

ADD ./apiserver /apiserver

ENTRYPOINT ["/apiserver"]
