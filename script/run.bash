#!/bin/bash
# 启动docker容器
cd ..

case "$1" in
  start)
    docker compose up -d
    ;;
  stop)
    docker compose down
    ;;
  *)
    echo "Usage: $0 {start|stop}"
    ;;
esac