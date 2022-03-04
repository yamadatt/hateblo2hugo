#FROM golang:1.17.7-alpine3.15
FROM golang:latest

# アップデートとgitのインストール！！
#RUN apk add --update &&  apk add git
RUN apt-get update && apt-get install git

# appディレクトリの作成
RUN mkdir /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app

# ホストのファイルをコンテナの作業ディレクトリに移行
#ADD . /go/src/app

#RUN go install github.com/akiradeveloper/hateblo2hugo@latest
RUN go get -u github.com/yamadatt/hateblo2hugo
