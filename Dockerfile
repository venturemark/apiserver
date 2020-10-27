FROM alpine:3.12

RUN apk add --no-cache ca-certificates

ADD ./apiserver /apiserver

ENTRYPOINT ["/apiserver"]
