FROM golang:1.13 as base

RUN apt update -y \
    && apt upgrade -y \
    && apt update -y \
    && apt install -y --no-install-recommends \
    less \
    && apt clean -y \
    && rm -rf /var/lib/apt/lists/*

# /workspace 以下はホストからボリュームをマウントする想定
# => コンテナ起動する度にvscode周り構築し直すのが手間なのでext系はここを使い回す
FROM base as godev

ENV HOME=/workspace
WORKDIR /workspace