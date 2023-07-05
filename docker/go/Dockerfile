FROM golang:1.20-alpine

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app/go-architecture-proto
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
