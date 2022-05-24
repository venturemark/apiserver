FROM alpine:3.16.0

RUN apk add --no-cache ca-certificates

ADD ./apiserver /apiserver
ADD ./ventures.json /ventures.json

ENTRYPOINT ["/apiserver"]
