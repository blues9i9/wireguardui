# WireGuard UI

一个用于管理 WireGuard 配置的 Web 用户界面。
支持中文、邮件集成页面配置、支持一件安装wireguard 

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

## 快速开始

> ⚠️ 默认用户名和密码为 `admin`，请登录后及时修改。

### 使用二进制文件

从 [Releases](https://github.com/ngoduykhanh/wireguard-ui/releases) 下载二进制文件并直接运行：

```bash
./wireguard-ui
```

### 使用 Docker Compose

参考 [examples/docker-compose](examples/docker-compose) 目录下的示例文件：

```bash
docker-compose up
```

## 环境变量

| 变量 | 说明 | 默认值 |
|------|------|--------|
| `BASE_PATH` | 反向代理子路径（如 `/wireguard`） | N/A |
| `BIND_ADDRESS` | 监听地址和端口 | `0.0.0.0:5000` |
| `SESSION_SECRET` | Session 加密密钥 | 随机生成 |
| `SESSION_SECRET_FILE` | Session 密钥文件路径 | N/A |
| `SESSION_MAX_DURATION` | 记住登录状态的有效天数 | `90` |
| `SUBNET_RANGES` | 地址子网划分范围 | N/A |
| `WGUI_USERNAME` | 初始管理员用户名 | `admin` |
| `WGUI_PASSWORD` | 初始管理员密码 | `admin` |
| `WGUI_FAVICON_FILE_PATH` | 自定义网站图标 | 内嵌 WireGuard 图标 |
| `WGUI_DNS` | 默认 DNS 服务器 | `1.1.1.1` |
| `WGUI_MTU` | 默认 MTU | `1450` |
| `WGUI_PERSISTENT_KEEPALIVE` | 默认持久保活间隔 | `15` |
| `WGUI_CONFIG_FILE_PATH` | WireGuard 配置文件路径 | `/etc/wireguard/wg0.conf` |
| `WGUI_LOG_LEVEL` | 日志级别：`DEBUG`、`INFO`、`WARN`、`ERROR`、`OFF` | `INFO` |
| `EMAIL_FROM_ADDRESS` | 发件邮箱地址 | N/A |
| `SENDGRID_API_KEY` | SendGrid API 密钥 | N/A |
| `SMTP_HOSTNAME` | SMTP 服务器地址 | `127.0.0.1` |
| `SMTP_PORT` | SMTP 端口 | `25` |
| `TELEGRAM_TOKEN` | Telegram 机器人 Token | N/A |

更多环境变量请参考原版文档。

## WireGuard 一键安装

在 **关于** 页面中提供了 WireGuard 自动检测与安装功能：

- 自动检测当前系统是否已安装 WireGuard 工具
- 支持多种包管理器：`apt`、`dnf`、`zypper`、`pacman`、`apk`
- 安装后自动加载内核模块并启用 IP 转发
- 已安装时显示"重新安装/升级"按钮

## 构建

### 构建二进制文件

```bash
# 编译 Linux AMD64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o wireguard-ui .
```

### 构建 Docker 镜像

```bash
docker build --build-arg=GIT_COMMIT=$(git rev-parse --short HEAD) -t wireguard-ui .
```

## 部署为系统服务

```bash
# 复制二进制文件
cp wireguard-ui /opt/wireguard-ui/

# 创建 systemd 服务（参考 deploy/wireguard-ui.service）
systemctl enable wireguard-ui
systemctl start wireguard-ui
```

## License

MIT. 参见 [LICENSE](https://github.com/ngoduykhanh/wireguard-ui/blob/master/LICENSE)。

## 致谢

- 原项目：[ngoduykhanh/wireguard-ui](https://github.com/ngoduykhanh/wireguard-ui)
- 二次改进者：blues
