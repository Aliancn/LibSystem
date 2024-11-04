#!/bin/zsh
# 启动docker容器
cd ..

if [ "$1" = "start" ]; then
docker compose up -d
elif [ "$1" = "stop" ]; then
  docker compose down
else
  echo "Usage: $0 {start|stop}"
fi

