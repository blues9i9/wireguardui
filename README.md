# WireGuard UI

一个用于管理 WireGuard 配置的 Web 用户界面。
支持中文、邮件集成、页面配置、WireGuard 一键安装。

![wireguard-ui](https://user-images.githubusercontent.com/37958026/177041280-e3e7ca16-d4cf-4e95-9920-68af15e780dd.png)

## 功能特点

- 友好的 Web 界面
- 用户认证（支持管理员/操作员角色）
- 多语言支持（中文/英文）
- 管理客户端信息（名称、邮箱等）
- 通过二维码/文件/邮件/Telegram 分发客户端配置
- 服务端配置管理（接口地址、监听端口、PostUp/PostDown 脚本）
- 全局设置管理（DNS、MTU、Keepalive、防火墙标记等）
- 客户端状态监控（实时查看在线/离线状态）
- Wake-on-LAN 远程唤醒
- **WireGuard 自动检测与一键安装**
- 会话管理与记住登录

## 快速开始（Linux 部署）

> ⚠️ 默认用户名和密码为 `admin`，请登录后及时修改。

### 1. 下载二进制文件

从 [Releases](https://github.com/blues9i9/wireguardui/releases) 下载对应平台的二进制文件：

```bash
# Linux AMD64
wget https://github.com/blues9i9/wireguardui/releases/download/v1.0.0/wireguard-ui-linux-amd64 -O /opt/wireguard-ui/wireguard-ui

# Linux ARM64（树莓派等）
wget https://github.com/blues9i9/wireguardui/releases/download/v1.0.0/wireguard-ui-linux-arm64 -O /opt/wireguard-ui/wireguard-ui

# Linux ARM（32位）
wget https://github.com/blues9i9/wireguardui/releases/download/v1.0.0/wireguard-ui-linux-arm -O /opt/wireguard-ui/wireguard-ui
```

### 2. 创建 systemd 服务

```bash
mkdir -p /opt/wireguard-ui
chmod +x /opt/wireguard-ui/wireguard-ui

cat > /etc/systemd/system/wireguard-ui.service << 'EOF'
[Unit]
Description=WireGuard UI
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/opt/wireguard-ui
ExecStart=/opt/wireguard-ui/wireguard-ui
Restart=on-failure
RestartSec=5
Environment=WGUI_USERNAME=admin
Environment=WGUI_PASSWORD=admin
Environment=WGUI_SESSION_SECRET=<替换为随机字符串>

[Install]
WantedBy=multi-user.target
EOF
```

> **重要**：将 `WGUI_SESSION_SECRET` 替换为随机字符串：`openssl rand -hex 32`

### 3. 启动服务

```bash
systemctl daemon-reload
systemctl enable wireguard-ui
systemctl start wireguard-ui
systemctl status wireguard-ui
```

### 4. 配置防火墙

```bash
# firewalld（CentOS/RHEL/Rocky Linux）
firewall-cmd --permanent --add-port=5000/tcp
firewall-cmd --reload

# ufw（Ubuntu/Debian）
ufw allow 5000/tcp

# iptables
iptables -A INPUT -p tcp --dport 5000 -j ACCEPT
```

### 5. 访问

打开浏览器访问 `http://服务器IP:5000`

## Docker 部署

```bash
docker run -d --name wireguard-ui \
  -p 5000:5000 \
  -e WGUI_USERNAME=admin \
  -e WGUI_PASSWORD=admin \
  -e WGUI_SESSION_SECRET=$(openssl rand -hex 32) \
  blues9i9/wireguardui:latest
```

或使用 docker-compose，参考 [examples/docker-compose](examples/docker-compose) 目录。

## 环境变量

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `BIND_ADDRESS` | 监听地址和端口 | `0.0.0.0:5000` |
| `BASE_PATH` | 反向代理子路径（如 `/wireguard`） | N/A |
| `SESSION_SECRET` | Session 加密密钥（**必改**） | 随机生成 |
| `SESSION_SECRET_FILE` | Session 密钥文件路径 | N/A |
| `SESSION_MAX_DURATION` | 记住登录状态的有效天数 | `90` |
| `SUBNET_RANGES` | 地址子网划分范围 | N/A |
| `WGUI_USERNAME` | 初始管理员用户名 | `admin` |
| `WGUI_PASSWORD` | 初始管理员密码 | `admin` |
| `WGUI_FAVICON_FILE_PATH` | 自定义网站图标 | 内嵌 WireGuard 图标 |
| `WGUI_DNS` | 默认 DNS 服务器 | `1.1.1.1` |
| `WGUI_MTU` | 默认 MTU | `1450` |
| `WGUI_PERSISTENT_KEEPALIVE` | 默认持久保活间隔 | `15` |
| `WGUI_ENDPOINT_ADDRESS` | 服务端公网地址 | 自动检测公网 IP |
| `WGUI_CONFIG_FILE_PATH` | WireGuard 配置文件路径 | `/etc/wireguard/wg0.conf` |
| `WGUI_SERVER_LISTEN_PORT` | WireGuard 监听端口 | `51820` |
| `WGUI_LOG_LEVEL` | 日志级别：`DEBUG`、`INFO`、`WARN`、`ERROR`、`OFF` | `INFO` |
| `WG_CONF_TEMPLATE` | 自定义 wg.conf 模板路径 | N/A |
| `EMAIL_FROM_ADDRESS` | 发件邮箱地址 | N/A |
| `EMAIL_FROM_NAME` | 发件人名称 | `WireGuard UI` |
| `SENDGRID_API_KEY` | SendGrid API 密钥 | N/A |
| `SMTP_HOSTNAME` | SMTP 服务器地址 | `127.0.0.1` |
| `SMTP_PORT` | SMTP 端口 | `25` |
| `SMTP_USERNAME` | SMTP 用户名 | N/A |
| `SMTP_PASSWORD` | SMTP 密码 | N/A |
| `SMTP_AUTH_TYPE` | SMTP 认证类型：`PLAIN`、`LOGIN`、`NONE` | `NONE` |
| `SMTP_ENCRYPTION` | 加密方式：`NONE`、`SSL`、`SSLTLS`、`TLS`、`STARTTLS` | `STARTTLS` |
| `TELEGRAM_TOKEN` | Telegram 机器人 Token | N/A |
| `TELEGRAM_ALLOW_CONF_REQUEST` | 允许用户通过 bot 获取配置 | `false` |

### 默认服务端配置

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `WGUI_SERVER_INTERFACE_ADDRESSES` | 服务端接口地址 | `10.252.1.0/24` |
| `WGUI_SERVER_LISTEN_PORT` | 服务端监听端口 | `51820` |
| `WGUI_SERVER_POST_UP_SCRIPT` | 启动后脚本 | N/A |
| `WGUI_SERVER_POST_DOWN_SCRIPT` | 停止后脚本 | N/A |

### 默认新客户端配置

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `WGUI_DEFAULT_CLIENT_ALLOWED_IPS` | 允许的 IP | `0.0.0.0/0` |
| `WGUI_DEFAULT_CLIENT_USE_SERVER_DNS` | 使用服务端 DNS | `true` |
| `WGUI_DEFAULT_CLIENT_ENABLE_AFTER_CREATION` | 创建后启用 | `true` |

### Docker 专用

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `WGUI_MANAGE_START` | 容器启停时启停 WireGuard | `false` |
| `WGUI_MANAGE_RESTART` | 配置变更后自动重启 WireGuard | `false` |

## WireGuard 一键安装

在 **关于** 页面提供了 WireGuard 自动检测与安装功能：

- 自动检测当前系统是否已安装 WireGuard 工具
- 支持多种包管理器：`apt`、`dnf`、`zypper`、`pacman`、`apk`
- 安装后自动加载内核模块并启用 IP 转发
- 已安装时显示"重新安装/升级"按钮

## 反向代理（Nginx）

### 根路径

```nginx
server {
    listen 80;
    server_name vpn.example.com;

    location / {
        proxy_pass http://127.0.0.1:5000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### 子路径（如 /wireguard）

```nginx
server {
    listen 80;
    server_name vpn.example.com;

    location /wireguard {
        proxy_pass http://127.0.0.1:5000;
    }
}
```

启动时需指定 `Environment=BASE_PATH=/wireguard`

### TLS/HTTPS（acme.sh + Let's Encrypt）

```bash
curl https://get.acme.sh | sh
~/.acme.sh/acme.sh --issue -d vpn.example.com --nginx
```

Nginx SSL 配置：

```nginx
server {
    listen 443 ssl;
    server_name vpn.example.com;

    ssl_certificate /root/.acme.sh/vpn.example.com/fullchain.cer;
    ssl_certificate_key /root/.acme.sh/vpn.example.com/vpn.example.com.key;

    location / {
        proxy_pass http://127.0.0.1:5000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

## WireGuard 自动重启配置

### systemd path 监控

```bash
cat > /etc/systemd/system/wgui.service << 'EOF'
[Unit]
Description=Restart WireGuard
After=network.target
[Service]
Type=oneshot
ExecStart=/usr/bin/systemctl restart wg-quick@wg0.service
[Install]
RequiredBy=wgui.path
EOF

cat > /etc/systemd/system/wgui.path << 'EOF'
[Unit]
Description=Watch /etc/wireguard/wg0.conf for changes
[Path]
PathModified=/etc/wireguard/wg0.conf
[Install]
WantedBy=multi-user.target
EOF

systemctl enable wgui.{path,service}
systemctl start wgui.{path,service}
```

### 全局设置 PostUp/Down

在 UI → 全局设置 → 配置：

```
PostUp = systemctl restart wg-quick@wg0
PostDown = systemctl stop wg-quick@wg0
```

## 构建

```bash
# 编译 Linux AMD64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o wireguard-ui .

# 编译所有平台
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o wireguard-ui-linux-amd64 .
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o wireguard-ui-linux-arm64 .
GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o wireguard-ui-linux-arm .
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o wireguard-ui-windows-amd64.exe .
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o wireguard-ui-darwin-amd64 .
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o wireguard-ui-darwin-arm64 .

# Docker 镜像
docker build --build-arg=GIT_COMMIT=$(git rev-parse --short HEAD) -t wireguard-ui .
```

## 日志查看

```bash
journalctl -u wireguard-ui -f     # 实时日志
journalctl -u wireguard-ui -n 50  # 最近50行
```

## 更新升级

```bash
wget https://github.com/blues9i9/wireguardui/releases/download/v新版本/wireguard-ui-linux-amd64 -O /opt/wireguard-ui/wireguard-ui
chmod +x /opt/wireguard-ui/wireguard-ui
systemctl restart wireguard-ui
```

## 常见问题

### 端口被占用

```bash
ss -tlnp | grep 5000
# 修改端口：在 service 中添加 Environment=BIND_ADDRESS=0.0.0.0:8080
```

### 无法发送邮件
检查 SMTP 配置或改用 SendGrid API。

### Session 失效
重启服务或重新登录。`SESSION_SECRET` 变更会导致所有会话失效。

## License

MIT. 参见 [LICENSE](https://github.com/blues9i9/wireguardui/blob/master/LICENSE)。

## 致谢

- 原项目：[ngoduykhanh/wireguard-ui](https://github.com/ngoduykhanh/wireguard-ui)
- 二次改进者：blues
