#!/bin/bash
set -e

REPO="Prominence673/tdocli"
BINARY_NAME="tdocli"
INSTALL_DIR="/usr/local/bin"

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
case "$OS" in
  linux*)  OS="linux" ;;
  darwin*) OS="darwin" ;;
  *) echo "Unsupported operating system: $OS"; exit 1 ;;
esac

ARCH=$(uname -m)
case "$ARCH" in
  x86_64)        ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

FILENAME="${BINARY_NAME}-${OS}-${ARCH}"
URL="https://github.com/${REPO}/releases/latest/download/${FILENAME}"

echo "Downloading ${FILENAME}..."
curl -sSL -o "/tmp/${BINARY_NAME}" "$URL"

chmod +x "/tmp/${BINARY_NAME}"

echo "Installing to ${INSTALL_DIR} (may prompt for your password)..."
sudo mv "/tmp/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"

echo "✓ ${BINARY_NAME} installed successfully"
"${INSTALL_DIR}/${BINARY_NAME}" --version