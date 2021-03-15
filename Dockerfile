##-----------------------stage-1------------------------##

FROM golang:1.16.2-alpine3.12 AS BUILD

WORKDIR /tmp/pxydot
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/pxydot && \
    chmod +x bin/pxydot

##-----------------------stage-2------------------------##

FROM alpine:3.12.4 AS RUNTIME

RUN apk --no-cache add ca-certificates
WORKDIR /opt/pxydot
COPY --from=BUILD /tmp/pxydot/bin/pxydot .
COPY --from=BUILD /tmp/pxydot/config.yaml .

RUN adduser --disabled-password --gecos '' pxydot
USER pxydot

EXPOSE 53/tcp
EXPOSE 53/udp

CMD ["/opt/pxydot/pxydot"]
