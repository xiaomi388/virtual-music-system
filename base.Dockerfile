FROM ubuntu:20.10

WORKDIR /

RUN apt-get update && \
    apt-get install -y golang-go \
    python3-pip git build-essential nodejs npm ffmpeg

# install netease api server
RUN git clone https://github.com/Binaryify/NeteaseCloudMusicApi.git && \
    cd ./NeteaseCloudMusicApi && npm install

# install youtube-dl
RUN pip3 install --upgrade youtube_dl

ENV LANG="C.UTF-8"