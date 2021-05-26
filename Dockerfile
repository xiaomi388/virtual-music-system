FROM xiaomi388/virtual-music-system-base:latest

COPY . /virtual-music-system
RUN cd ./virtual-music-system/ && GO111MODULE=on go build

EXPOSE 3001

ENV VMS_HTTP_HOSTPORT="0.0.0.0:3001"
ENV VMS_METADATA_DRIVER_NETEASE_BASEURL="http://localhost:3000"
ENV VMS_SONG_LOCALREPO_FILE_ROOTDIR="/data"

WORKDIR /virtual-music-system/scripts

CMD ./docker_run.sh