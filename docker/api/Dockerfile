ARG GO_VERSION=${GO_VERSION}
ARG API_CODE_WORKDIR=${API_CODE_WORKDIR}

FROM golang:${GO_VERSION}-alpine

RUN apk update && apk add git

# ワーキングディレクトリの設定
WORKDIR /src/api