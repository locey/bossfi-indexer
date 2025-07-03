#!/bin/bash

# 获取当前脚本所在目录（项目根目录）
PROJECT_DIR="$(cd "$(dirname "$0")" && pwd)"
IMAGE_NAME="bossfi-indexer"
PORT=8000

# 用法提示
usage() {
  echo "用法: ./deploy.sh v1.0.0"
  exit 1
}

# 校验版本号参数
if [ $# -ne 1 ]; then
  echo "❌ 错误：必须传入版本号参数"
  usage
fi

VERSION=$1

# 校验版本号格式（v1.0.0 形式）
if [[ ! "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "❌ 错误：版本号格式不正确，必须是 vX.Y.Z，例如 v1.0.1"
  usage
fi

echo "========== 🚀 部署版本：$VERSION =========="

cd "$PROJECT_DIR" || { echo "❌ 无法进入项目目录 $PROJECT_DIR"; exit 1; }

# 拉取指定 Git tag
echo "1️⃣ 拉取 Git tag：$VERSION ..."
git fetch --tags
git checkout "$VERSION" || { echo "❌ 找不到对应的 Git tag：$VERSION"; exit 1; }

# 构建镜像
echo "2️⃣ 构建镜像：$IMAGE_NAME:$VERSION ..."
docker build -t "$IMAGE_NAME:$VERSION" .

# 停止并删除旧容器
echo "3️⃣ 停止并删除旧容器（如果存在）..."
docker stop "$IMAGE_NAME" >/dev/null 2>&1 || true
docker rm "$IMAGE_NAME" >/dev/null 2>&1 || true

# 启动新容器
echo "4️⃣ 启动容器..."
docker run -d \
  --name "$IMAGE_NAME" \
  -p $PORT:$PORT \
  -v /opt/bossfi/bossfi-indexer-configs/config.toml:/data/configs/config.toml \
  --restart=always \
  "$IMAGE_NAME:$VERSION"

echo "$(date +"%Y-%m-%d %H:%M:%S") deployed $PORT $VERSION" >> logs/deployed.log
echo "✅ 部署成功：$IMAGE_NAME:$VERSION 正在运行于端口 $PORT"
