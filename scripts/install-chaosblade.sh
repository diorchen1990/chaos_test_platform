#!/bin/bash

set -e

# ChaosBlade版本
CHAOSBLADE_VERSION="1.7.0"
INSTALL_DIR="/opt/chaosblade"

# 下载ChaosBlade
echo "下载 ChaosBlade..."
wget https://github.com/chaosblade-io/chaosblade/releases/download/v${CHAOSBLADE_VERSION}/chaosblade-${CHAOSBLADE_VERSION}-linux-amd64.tar.gz

# 解压安装
echo "安装 ChaosBlade..."
tar -xvf chaosblade-${CHAOSBLADE_VERSION}-linux-amd64.tar.gz
mv chaosblade-${CHAOSBLADE_VERSION} ${INSTALL_DIR}

# 创建软链接
ln -sf ${INSTALL_DIR}/bin/blade /usr/local/bin/blade

# 验证安装
blade version

# 安装必要的依赖
echo "安装依赖..."
blade prepare --install-dependency

# 安装Java Agent(如果需要)
echo "安装Java Agent..."
blade prepare jvm 