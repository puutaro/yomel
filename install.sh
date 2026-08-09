#!/bin/bash
set -eu

readonly CMD_NAME=yomel
readonly REPO="puutaro/${CMD_NAME}"

# OSの判定 (mac/linux共通)
readonly OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$OS" in
  linux|darwin) ;;
  *) echo "Error: Unsupported OS: $OS"; exit 1 ;;
esac

# アーキテクチャの判定
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Error: Unsupported architecture: $ARCH"; exit 1 ;;
esac

# GitHub Releasesから最新のバージョンタグ（例: v0.0.3）を自動取得
echo "==> Fetching the latest version..."
LATEST_VERSION=$(curl -sI "https://github.com/${REPO}/releases/latest" | grep -i "^location:" | awk -F'/' '{print $NF}' | tr -d '\r')

if [ -z "$LATEST_VERSION" ]; then
  echo "Error: Failed to fetch the latest version."
  exit 1
fi

readonly DOWNLOAD_FILENAME="${CMD_NAME}_${LATEST_VERSION}_${OS}_${ARCH}"
readonly URL="https://github.com/${REPO}/releases/download/${LATEST_VERSION}/${DOWNLOAD_FILENAME}"

echo "==> Downloading ${CMD_NAME} ${LATEST_VERSION} for ${OS}/${ARCH}..."
curl -sL -o "${DOWNLOAD_FILENAME}" "${URL}"

echo "==> Installing to /usr/local/bin/${CMD_NAME}..."
chmod +x "${DOWNLOAD_FILENAME}"
sudo mv "${DOWNLOAD_FILENAME}" "/usr/local/bin/${CMD_NAME}"

echo "==> Successfully installed! You can now run: ${CMD_NAME} --help"