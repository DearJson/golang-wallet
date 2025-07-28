# 🚀 钱包管理系统 Docker 部署指南

## 📋 系统要求

- Docker >= 20.0
- Docker Compose >= 2.0  
- 系统内存 >= 4GB
- 磁盘空间 >= 10GB

## ⚡ 快速开始

### 1. 克隆项目
```bash
git clone <your-repo-url>
cd golang-wallet
```

### 2. 配置环境变量
```bash
# 编辑配置文件，修改默认密码
vim docker.env

# 必须修改的配置项：
# - MYSQL_ROOT_PASSWORD
# - MYSQL_PASSWORD  
# - REDIS_PASSWORD
# - RABBITMQ_PASSWORD
```

### 3. 一键部署
```bash
# 构建并启动所有服务
./build.sh

# 或者自动生成随机密码
./build.sh build
```

### 4. 验证部署
```bash
# 检查服务状态
./build.sh status

# 查看日志
./build.sh logs
```

## 🎯 访问地址

| 服务 | 地址 | 说明 |
|------|------|------|
| 钱包管理系统 | http://localhost:9090 | 主应用 |
| RabbitMQ管理 | http://localhost:15672 | 消息队列管理界面 |
| MySQL | localhost:3307 | 数据库 |
| Redis | localhost:6379 | 缓存 |

## 🛠️ 常用命令

### 服务管理
```bash
# 启动服务
./build.sh start

# 停止服务  
./build.sh stop

# 重启服务
./build.sh restart

# 查看状态
./build.sh status

# 查看日志
./build.sh logs

# 清理环境
./build.sh clean
```

### Docker命令
```bash
# 查看容器状态
docker-compose ps

# 进入容器
docker-compose exec wallet-app sh
docker-compose exec mysql bash
docker-compose exec redis sh

# 查看资源使用
docker stats

# 备份数据
docker-compose exec mysql mysqldump -u root -p wallet > backup.sql
```

## 🔧 配置修改

### 应用配置
主要配置文件：`config/config.toml`

```toml
# 修改数据库连接
[database]
link = "mysql:wallet:wallet123@tcp(mysql:3306)/wallet"

# 修改Redis连接  
[redis]
default = "redis:6379,0?idleTimeout=20&maxActive=100"

# 修改RabbitMQ连接
[rabbitmq]
host = "rabbitmq"
port = "5672"
user = "admin"
password = "admin123"
```

### 端口映射修改
编辑 `docker-compose.yml`：
```yaml
ports:
  - "8080:9090"  # 改为8080端口
```

## 📊 监控和维护

### 健康检查
```bash
# 应用健康检查
curl http://localhost:9090/health

# 数据库连接检查
docker-compose exec mysql mysqladmin ping

# Redis连接检查  
docker-compose exec redis redis-cli ping
```

### 日志管理
```bash
# 实时查看应用日志
docker-compose logs -f wallet-app

# 查看错误日志
docker-compose logs wallet-app | grep ERROR

# 清理日志（重启容器）
docker-compose restart wallet-app
```

### 数据备份
```bash
# 数据库备份
docker-compose exec mysql mysqldump -u root -p$(grep MYSQL_ROOT_PASSWORD docker.env | cut -d= -f2) wallet > backup_$(date +%Y%m%d).sql

# Redis备份
docker-compose exec redis redis-cli BGSAVE

# 文件备份
tar -czf data_backup_$(date +%Y%m%d).tar.gz data/
```

## 🚨 故障排除

### 常见问题

#### 1. 端口占用
```bash
# 查看端口使用情况
netstat -tlnp | grep :9090

# 停止占用端口的进程
sudo kill -9 <PID>
```

#### 2. 权限问题
```bash
# 修复目录权限
sudo chown -R $(whoami):$(whoami) data/ logs/
chmod 755 data/ logs/
```

#### 3. 内存不足
```bash
# 清理Docker缓存
docker system prune -f

# 清理未使用的镜像
docker image prune -f
```

#### 4. 数据库连接失败
```bash
# 检查MySQL容器状态
docker-compose logs mysql

# 重启MySQL
docker-compose restart mysql

# 检查数据库配置
docker-compose exec mysql mysql -u root -p -e "SHOW DATABASES;"
```

### 调试模式
```bash
# 前台运行查看详细日志
docker-compose up

# 单独启动应用容器进行调试
docker-compose run --rm wallet-app sh
```

## 🔒 安全配置

### 生产环境建议
1. **修改默认端口**：避免使用默认端口
2. **使用HTTPS**：配置SSL证书
3. **限制网络访问**：配置防火墙规则
4. **定期更新**：更新基础镜像和依赖
5. **监控日志**：设置日志告警

### 密码策略
- 使用至少16位的随机密码
- 包含大小写字母、数字、特殊字符
- 定期更换密码（建议3个月）
- 不要在配置文件中使用明文密码

## 📝 更新和维护

### 应用更新
```bash
# 拉取最新代码
git pull

# 重新构建
./build.sh clean
./build.sh build
```

### 镜像更新
```bash
# 更新基础镜像
docker-compose pull

# 重新构建应用镜像
docker-compose build --no-cache wallet-app
```

### 数据迁移
```bash
# 导出数据
docker-compose exec mysql mysqldump -u root -p wallet > migration.sql

# 导入数据
docker-compose exec -T mysql mysql -u root -p wallet < migration.sql
```

## 💡 最佳实践

1. **定期备份**：每日自动备份重要数据
2. **监控资源**：监控CPU、内存、磁盘使用情况
3. **日志轮转**：配置日志轮转避免磁盘满
4. **安全更新**：及时更新系统和应用补丁
5. **容量规划**：根据业务增长调整资源配置

## 📞 技术支持

如遇到问题，请提供以下信息：
- 系统版本：`uname -a`
- Docker版本：`docker --version`
- 容器状态：`docker-compose ps`
- 错误日志：`docker-compose logs`

---

🎉 **恭喜！你的钱包管理系统已成功部署！** 