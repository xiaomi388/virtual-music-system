#!/bin/bash

cd /NeteaseCloudMusicApi/ && node app.js &
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start netease api: $status"
  exit $status
fi

cd /virtual-music-system/ && go test ./...
status=$?
if [ $status -ne 0 ]; then
  echo "Test Failed: $status"
  exit $status
fi

