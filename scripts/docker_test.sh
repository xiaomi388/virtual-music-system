#!/bin/bash

SCRIPT_DIR=$(dirname "$0")

cd /NeteaseCloudMusicApi/ && node app.js &
status=$?
if [ $status -ne 0 ]; then
  echo "Failed to start netease api: $status"
  exit $status
fi

cd "$SCRIPT_DIR"/.. && go test ./...
status=$?
if [ $status -ne 0 ]; then
  echo "Test Failed: $status"
  exit $status
fi

