#!/bin/bash

# 获取版本号
# 优先使用 git tag，如果没有 tag 则使用 git commit 的短 hash
VERSION=$(git describe --tags 2>/dev/null || git rev-parse --short HEAD)

# 编译 Linux AMD64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/nacos-cli_linux_amd64

# 编译 Linux ARM64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o dist/nacos-cli_linux_arm64

# 编译 Windows AMD64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/nacos-cli_windows_amd64.exe

# 添加可执行权限
mkdir -p dist
chmod +x dist/nacos-cli_linux_*

# 打包发布文件
cd dist
tar -czf nacos-cli_${VERSION}_linux_amd64.tar.gz nacos-cli_linux_amd64
tar -czf nacos-cli_${VERSION}_linux_arm64.tar.gz nacos-cli_linux_arm64
zip nacos-cli_${VERSION}_windows_amd64.zip nacos-cli_windows_amd64.exe