#!/bin/bash

readonly CMD_NAME=yomel
readonly YOME_VERSION="0.0.1"

readonly OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$OS" in
  linux|darwin) ;;
  *) echo "Error: Unsupported OS: $OS"; exit 1 ;;
esac

ARCH="$(uname -m)"
case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Error: Unsupported architecture: $ARCH"; exit 1 ;;
esac

readonly DOWNLOAD_FILENAME="${CMD_NAME}_v${YOME_VERSION}_${OS}_${ARCH}"
readonly URL="https://github.com/puutaro/${CMD_NAME}/releases/download/v${YOME_VERSION}/${DOWNLOAD_FILENAME}"

echo "==> Downloading ${CMD_NAME} v${YOME_VERSION} for ${OS}/${ARCH}..."
curl -L -o "${DOWNLOAD_FILENAME}" "${URL}"

echo "==> Installing to /usr/local/bin/${CMD_NAME}..."
chmod +x "${DOWNLOAD_FILENAME}"
sudo mv "${DOWNLOAD_FILENAME}" /usr/local/bin/${CMD_NAME}

echo "==> Successfully installed! You can now run: ${CMD_NAME} --help"
