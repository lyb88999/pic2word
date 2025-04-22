#!/bin/bash

# Pic2Word 部署脚本
# 作者: Claude
# 日期: 2024年4月20日

set -e  # 遇到错误立即退出

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
NC='\033[0m' # 无颜色

# 配置变量
REPO_URL="https://github.com/lyb88999/pic2word.git"
INSTALL_DIR="/opt/pic2word"
FRONTEND_PORT=80
BACKEND_PORT=8080
DOMAIN=""  # 替换为您的域名
GOOGLE_API_KEY="AIzaSyAbL3mNHTnk8_FlOG8S978UKQy274OC-MQ"  # 替换为您的Google API密钥

# 函数：打印带颜色的信息
print_message() {
  echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
  echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
  echo -e "${RED}[ERROR]${NC} $1"
}

# 函数：检查命令是否存在
check_command() {
  if ! command -v $1 &> /dev/null; then
    print_error "$1 未安装。正在尝试安装..."
    return 1
  else
    return 0
  fi
}

# 确保脚本以root权限运行
if [ "$EUID" -ne 0 ]; then
  print_error "请以root权限运行此脚本"
  exit 1
fi

# 获取用户输入
read -p "请输入您的Google API密钥: " GOOGLE_API_KEY
read -p "请输入您的域名 (不填则使用IP访问): " DOMAIN_INPUT
if [ ! -z "$DOMAIN_INPUT" ]; then
  DOMAIN=$DOMAIN_INPUT
fi

print_message "开始部署 Pic2Word 应用..."
print_message "安装目录: $INSTALL_DIR"
print_message "域名: $DOMAIN"

# 更新系统包
print_message "正在更新系统包..."
apt-get update
apt-get upgrade -y

# 安装基本依赖
print_message "正在安装基本依赖..."
apt-get install -y git curl wget build-essential

# 安装Go
if ! check_command go; then
  print_message "正在安装Go..."
  wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
  tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
  echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
  source /etc/profile
  rm go1.21.0.linux-amd64.tar.gz
  print_message "Go 安装完成"
else
  print_message "Go 已安装，版本: $(go version)"
fi

# 安装Node.js和npm
if ! check_command node; then
  print_message "正在安装Node.js和npm..."
  curl -fsSL https://deb.nodesource.com/setup_20.x | bash -
  apt-get install -y nodejs
  print_message "Node.js 安装完成"
else
  print_message "Node.js 已安装，版本: $(node -v)"
fi

# 安装Pandoc
if ! check_command pandoc; then
  print_message "正在安装Pandoc..."
  apt-get install -y pandoc
  print_message "Pandoc 安装完成"
else
  print_message "Pandoc 已安装，版本: $(pandoc --version | head -n 1)"
fi

# 安装Nginx
if ! check_command nginx; then
  print_message "正在安装Nginx..."
  apt-get install -y nginx
  print_message "Nginx 安装完成"
else
  print_message "Nginx 已安装，版本: $(nginx -v 2>&1 | cut -d '/' -f 2)"
fi

# 克隆项目
print_message "正在克隆项目仓库..."
if [ -d "$INSTALL_DIR" ]; then
  print_warning "目录 $INSTALL_DIR 已存在，更新代码..."
  cd $INSTALL_DIR
  git pull
else
  git clone $REPO_URL $INSTALL_DIR
  cd $INSTALL_DIR
fi

# 创建环境配置文件
print_message "正在创建环境配置文件..."

# 后端配置
cat > $INSTALL_DIR/backend/.env << EOF
# 服务器配置
PORT=$BACKEND_PORT
ENV=production

# Google Gemini API配置
GOOGLE_API_KEY='$GOOGLE_API_KEY'
GOOGLE_PROJECT_ID=''
GOOGLE_LOCATION=us-central1

# 临时文件目录
TEMP_DIR=./tmp

# CORS配置
ALLOW_ORIGINS=http://$DOMAIN
EOF

# 前端配置
cat > $INSTALL_DIR/pic2word-frontend/.env << EOF
VITE_API_URL=http://$DOMAIN/api
EOF

# 构建后端
print_message "正在构建后端..."
cd $INSTALL_DIR/backend
mkdir -p tmp
go mod download
go build -o pic2word main.go

# 构建前端
print_message "正在构建前端..."
cd $INSTALL_DIR/pic2word-frontend
npm install
npm run build

# 创建系统服务
print_message "正在创建系统服务..."
cat > /etc/systemd/system/pic2word.service << EOF
[Unit]
Description=Pic2Word Backend Service
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=$INSTALL_DIR/backend
ExecStart=$INSTALL_DIR/backend/pic2word
Restart=on-failure
RestartSec=5
StartLimitInterval=60
StartLimitBurst=3

[Install]
WantedBy=multi-user.target
EOF

# 重新加载systemd并启动服务
systemctl daemon-reload
systemctl enable pic2word
systemctl start pic2word

# 配置Nginx
print_message "正在配置Nginx..."
cat > /etc/nginx/sites-available/pic2word << EOF
server {
    listen 80;
    server_name $DOMAIN;

    # 前端静态文件
    location / {
        root $INSTALL_DIR/pic2word-frontend/dist;
        index index.html;
        try_files \$uri \$uri/ /index.html;
    }

    # 后端API代理
    location /api/ {
        proxy_pass http://localhost:$BACKEND_PORT/api/;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }

    # 健康检查
    location /health {
        proxy_pass http://localhost:$BACKEND_PORT/health;
    }
}
EOF

# 启用Nginx配置
ln -sf /etc/nginx/sites-available/pic2word /etc/nginx/sites-enabled/
rm -f /etc/nginx/sites-enabled/default
nginx -t && systemctl restart nginx

print_message "==============================================="
print_message "Pic2Word 部署完成！"
print_message "前端访问地址: http://$DOMAIN"
print_message "后端API地址: http://$DOMAIN/api"
print_message "==============================================="
print_message "如需查看后端日志，请使用: journalctl -u pic2word"
print_message "重启服务命令: systemctl restart pic2word"
print_message "重启Nginx命令: systemctl restart nginx" 