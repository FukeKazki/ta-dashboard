FROM golang:1.20-alpine as development

ENV LANG C.UTF-8
ENV APP_ROOT /var/app

WORKDIR ${APP_ROOT}

RUN apk update && \
  apk upgrade --no-cache libcrypto3 libssl3

RUN apk add gcc g++

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# sql-migrateのインストール
# https://github.com/rubenv/sql-migrate/pull/259
RUN CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go install github.com/rubenv/sql-migrate/...@v1.6.0

# 開発時のみ必要なツール
RUN go install github.com/cosmtrek/air@v1.40.4 && \
  go install github.com/swaggo/swag/cmd/swag@v1.16.2 && \
  go install github.com/securego/gosec/v2/cmd/gosec@v2.14.0 && \
  go install honnef.co/go/tools/cmd/staticcheck@v0.4.3

CMD ["air", "-c", ".air.toml"]
