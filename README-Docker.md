# 🐳 钱包管理系统 Docker 部署方案

[![Docker](https://img.shields.io/badge/Docker-20.10+-blue.svg)](https://www.docker.com/)
[![Go](https://img.shields.io/badge/Go-1.19+-00ADD8.svg)](https://golang.org/)
[![Security](https://img.shields.io/badge/Security-Hardened-green.svg)](./docker-security.md)

## 🎯 项目概述

本项目为**钱包管理系统**的完整Docker化部署方案，实现了：

- 🔒 **高度安全**的"锁死"容器配置
- 🚀 **一键部署**的自动化脚本
- 📊 **生产级别**的性能优化
- 🛡️ **多重防护**的安全措施

## 📋 系统架构

```
┌─────────────────────────────────────────────────────────┐
│                    Docker 容器集群                        │
├─────────────────────────────────────────────────────────┤
│  Nginx (反向代理)  │  钱包应用 (Go)  │  监控服务           │
├─────────────────────────────────────────────────────────┤
│  MySQL (数据库)    │  Redis (缓存)   │  RabbitMQ (队列)   │
├─────────────────────────────────────────────────────────┤
│              Docker 网络隔离 & 安全加固                   │
└─────────────────────────────────────────────────────────┘
```

## 🔥 核心特性

### 🔒 安全加固（"锁死"容器）
- ❌ **禁止特权提升**：`no-new-privileges:true`
- 🛡️ **移除危险权限**：只保留必要的 Linux capabilities
- 👤 **非root运行**：所有服务以非特权用户运行
- 🌐 **网络隔离**：独立的Docker网络段
- 📁 **文件系统保护**：只读配置文件、临时文件系统

### ⚡ 性能优化
- 🏗️ **多阶段构建**：最小化镜像体积（<50MB）
- 📦 **静态编译**：零依赖的二进制文件
- 🔄 **连接池优化**：数据库和Redis连接池配置
- 📊 **资源限制**：CPU和内存使用限制

### 🚀 部署便利
- 🎯 **一键部署**：`./build.sh` 即可完成所有操作
- 🔐 **自动密码生成**：可选的随机强密码生成
- 📋 **健康检查**：自动检测服务状态
- 📝 **详细日志**：结构化日志记录

## 📁 文件结构

```
golang-wallet/
├── 🐳 Docker相关
│   ├── Dockerfile                    # 应用镜像构建
│   ├── docker-compose.yml            # 开发环境配置
│   ├── docker-compose.production.yml # 生产环境配置
│   ├── .dockerignore                 # Docker忽略文件
│   └── build.sh                      # 自动化构建脚本
├── ⚙️ 配置文件
│   ├── docker.env                    # 环境变量配置
│   ├── config/docker.toml           # Docker专用配置
│   └── config/config.toml           # 原始配置文件
├── 📚 文档
│   ├── Docker部署指南.md             # 详细部署指南
│   ├── docker-security.md           # 安全配置说明
│   └── README-Docker.md             # 本文件
└── 💾 数据目录
    ├── data/                         # 应用数据
    ├── logs/                         # 日志文件
    └── sql/                          # 数据库初始化脚本
```

## 🚀 快速开始

### 1️⃣ 环境准备
```bash
# 确保已安装Docker和Docker Compose
docker --version
docker-compose --version

# 克隆项目
git clone <your-repo-url>
cd golang-wallet
```

### 2️⃣ 配置安全密码
```bash
# 编辑环境变量文件
vim docker.env

# 修改以下密码（必须）：
# - MYSQL_ROOT_PASSWORD
# - MYSQL_PASSWORD
# - REDIS_PASSWORD
# - RABBITMQ_PASSWORD
```

### 3️⃣ 一键部署
```bash
# 构建并启动所有服务
./build.sh

# 或者让系统自动生成随机密码
./build.sh build
```

### 4️⃣ 验证部署
```bash
# 检查服务状态
./build.sh status

# 访问应用
curl http://localhost:9090
```

## 🎮 命令参考

| 命令 | 说明 | 用途 |
|------|------|------|
| `./build.sh` | 完整构建部署 | 首次部署 |
| `./build.sh start` | 启动服务 | 日常启动 |
| `./build.sh stop` | 停止服务 | 维护时停止 |
| `./build.sh restart` | 重启服务 | 配置更新后 |
| `./build.sh logs` | 查看日志 | 问题排查 |
| `./build.sh status` | 查看状态 | 健康检查 |
| `./build.sh clean` | 清理环境 | 重新部署前 |

## 🌍 环境配置

### 开发环境
```bash
# 使用默认配置
docker-compose up -d
```

### 生产环境
```bash
# 使用生产配置
docker-compose -f docker-compose.production.yml up -d
```

## 🔐 安全特性详解

### 容器权限限制
- **无特权容器**：禁止容器获取新权限
- **能力移除**：移除所有不必要的Linux capabilities
- **AppArmor防护**：启用应用程序防护
- **Seccomp过滤**：系统调用过滤

### 网络安全
- **网络隔离**：独立的Docker网络
- **端口限制**：只暴露必要端口
- **内部通信**：服务间使用内部网络

### 文件系统安全
- **只读根文件系统**：防止恶意文件写入
- **临时文件系统**：`/tmp`目录限制
- **权限最小化**：文件权限最小化原则

### 资源限制
- **CPU限制**：防止CPU过度占用
- **内存限制**：防止内存泄漏
- **进程数限制**：防止fork炸弹攻击

## 📊 性能指标

| 项目 | 指标 | 说明 |
|------|------|------|
| 🖼️ 镜像大小 | ~45MB | 多阶段构建优化 |
| 🚀 启动时间 | <30s | 包含所有依赖服务 |
| 💾 内存占用 | ~256MB | 应用运行时占用 |
| 🔄 并发连接 | 1000+ | 优化的连接池配置 |

## 🔍 监控和运维

### 健康检查
```bash
# 应用健康检查
curl http://localhost:9090/health

# 容器健康状态
docker-compose ps
```

### 日志管理
```bash
# 查看实时日志
docker-compose logs -f wallet-app

# 查看错误日志
docker-compose logs wallet-app | grep ERROR
```

### 数据备份
```bash
# 数据库备份
docker-compose exec mysql mysqldump -u root -p wallet > backup.sql

# 文件数据备份
tar -czf data_backup.tar.gz data/
```

## 🚨 故障排除

### 常见问题及解决方案

#### Port already in use
```bash
# 查看端口占用
netstat -tlnp | grep :9090
# 停止占用进程或修改端口配置
```

#### Permission denied
```bash
# 修复权限问题
sudo chown -R $(whoami):$(whoami) data/ logs/
chmod 755 data/ logs/
```

#### Container fails to start
```bash
# 查看详细错误日志
docker-compose logs <service-name>
# 检查配置文件语法
```

## 📈 性能调优建议

### 生产环境优化
1. **增加资源配置**：根据负载调整CPU和内存限制
2. **数据库优化**：调整MySQL配置参数
3. **缓存优化**：配置Redis持久化策略
4. **监控配置**：添加Prometheus + Grafana监控

### 高可用部署
1. **负载均衡**：使用Nginx进行负载均衡
2. **数据库集群**：配置MySQL主从复制
3. **缓存集群**：配置Redis集群模式
4. **容器编排**：使用Kubernetes进行容器编排

## 🤝 贡献指南

欢迎提交Issue和Pull Request来改进这个Docker部署方案！

### 开发流程
1. Fork本项目
2. 创建功能分支
3. 提交更改
4. 创建Pull Request

## 📜 许可证

本项目采用 [MIT License](LICENSE) 许可证。

## 📞 技术支持

- 📧 **邮箱支持**：support@yourcompany.com
- 💬 **在线交流**：加入我们的技术群
- 📖 **详细文档**：[完整部署指南](./Docker部署指南.md)
- 🔒 **安全说明**：[安全配置文档](./docker-security.md)

---

<div align="center">

**🎉 感谢使用钱包管理系统Docker部署方案！**

如果这个项目对你有帮助，请给我们一个 ⭐ Star！

[报告Bug](https://github.com/yourrepo/issues) • [功能建议](https://github.com/yourrepo/issues) • [加入讨论](https://github.com/yourrepo/discussions)

</div> 