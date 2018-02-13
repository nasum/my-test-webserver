FROM alpine:3.7

ENV APP_ROOT /app

WORKDIR $APP_ROOT

RUN apk update && apk add ca-certificates

COPY my-test-webserver $APP_ROOT/

EXPOSE 1323

CMD ./my-test-webserver

