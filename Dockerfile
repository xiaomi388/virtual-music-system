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


# install vms
COPY . /virtual-music-system
RUN cd ./virtual-music-system/ && GO111MODULE=on go build

EXPOSE 3001

ENV VMS_HTTP_HOSTPORT="0.0.0.0:3001"
ENV VMS_METADATA_DRIVER_NETEASE_BASEURL="http://localhost:3000"
ENV VMS_LOCALREPO_FILE_ROOTDIR="/data"

WORKDIR /virtual-music-system/scripts

CMD ./docker_run.sh