FROM alpine:3.15.2

RUN apk add --no-cache ca-certificates

ADD ./apiserver /apiserver
ADD ./ventures.json /ventures.json

ENTRYPOINT ["/apiserver"]
