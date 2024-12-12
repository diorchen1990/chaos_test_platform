#!/bin/bash

set -e

# 默认值
AGENT_TYPE=""
APP_NAME=""
SERVER_URL=""
CONFIG_FILE="agent-config.yaml"

# 解析参数
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        --agent-type)
            AGENT_TYPE="$2"
            shift
            shift
            ;;
        --app-name)
            APP_NAME="$2"
            shift
            shift
            ;;
        --server-url)
            SERVER_URL="$2"
            shift
            shift
            ;;
        *)
            echo "未知参数: $1"
            exit 1
            ;;
    esac
done

# 验证必要参数
if [ -z "$AGENT_TYPE" ] || [ -z "$APP_NAME" ] || [ -z "$SERVER_URL" ]; then
    echo "缺少必要参数"
    echo "用法: $0 --agent-type <type> --app-name <name> --server-url <url>"
    exit 1
fi

# 创建工作目录
WORK_DIR="/opt/chaos/agents/$APP_NAME"
mkdir -p "$WORK_DIR"
cd "$WORK_DIR"

# 下载对应的Agent
echo "下载 $AGENT_TYPE Agent..."
case $AGENT_TYPE in
    java)
        curl -L -o chaos-agent.jar "$SERVER_URL/download/agents/java/latest"
        ;;
    golang)
        curl -L -o chaos-agent.tar.gz "$SERVER_URL/download/agents/golang/latest"
        tar xzf chaos-agent.tar.gz
        ;;
    nodejs)
        curl -L -o chaos-agent.tgz "$SERVER_URL/download/agents/nodejs/latest"
        tar xzf chaos-agent.tgz
        ;;
    *)
        echo "不支持的Agent类型: $AGENT_TYPE"
        exit 1
        ;;
esac

# 生成配置文件
cat > $CONFIG_FILE << EOF
agent:
  name: $APP_NAME
  type: $AGENT_TYPE
  server: $SERVER_URL
  
metrics:
  enable: true
  interval: 15

logs:
  level: info
  path: /var/log/chaos-agent.log
EOF

# 安装系统服务
cat > /etc/systemd/system/chaos-agent.service << EOF
[Unit]
Description=Chaos Platform Agent
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=$WORK_DIR
ExecStart=/bin/bash -c 'case "$AGENT_TYPE" in \
    java) java -javaagent:chaos-agent.jar -jar chaos-agent.jar ;; \
    golang) ./chaos-agent ;; \
    nodejs) node chaos-agent.js ;; \
esac'
Restart=always

[Install]
WantedBy=multi-user.target
EOF

# 启动服务
systemctl daemon-reload
systemctl enable chaos-agent
systemctl start chaos-agent

echo "Agent安装完成！"
echo "查看状态: systemctl status chaos-agent"
echo "查看日志: journalctl -u chaos-agent -f" 