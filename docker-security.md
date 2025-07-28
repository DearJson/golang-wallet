# Docker 安全配置说明

## 🔒 "锁死"容器的安全措施

本项目实现了多重安全措施来"锁死"Docker容器，确保系统安全：

### 1. 容器权限限制

```yaml
security_opt:
  - no-new-privileges:true  # 禁止容器内进程获取新的权限
  - apparmor:docker-default # 启用AppArmor安全模块

cap_drop:
  - ALL  # 移除所有Linux capabilities

cap_add:
  - NET_BIND_SERVICE  # 只允许绑定网络端口
```

### 2. 文件系统安全

```yaml
# 临时文件系统（防止恶意文件写入）
tmpfs:
  - /tmp:size=100M,noexec,nosuid,nodev

# 只读文件系统（可选，当前因日志需求暂时关闭）
read_only: false
```

### 3. 资源限制

```yaml
deploy:
  resources:
    limits:
      cpus: '2.0'      # CPU限制
      memory: 1G       # 内存限制
    reservations:
      cpus: '0.5'      # CPU预留
      memory: 256M     # 内存预留
```

### 4. 网络隔离

```yaml
networks:
  wallet-network:
    driver: bridge
    ipam:
      config:
        - subnet: 172.20.0.0/16  # 独立网络段
```

### 5. 用户权限

- 🚫 **非root用户**：所有容器都以非特权用户运行
- 🔐 **UID/GID隔离**：使用固定的用户ID避免权限提升
- 📁 **最小权限原则**：只给予必要的文件访问权限

### 6. 镜像安全

- 🏗️ **多阶段构建**：最小化运行时镜像体积
- 🧹 **无调试工具**：生产镜像不包含shell、编译器等工具
- 📦 **静态编译**：减少运行时依赖

## 🛡️ 额外安全建议

### 防火墙配置
```bash
# 只开放必要端口
sudo ufw allow 9090/tcp   # 应用端口
sudo ufw allow 3307/tcp   # MySQL端口（如需外部访问）
sudo ufw deny 6379/tcp    # Redis（禁止外部访问）
sudo ufw deny 5672/tcp    # RabbitMQ（禁止外部访问）
```

### 日志监控
```bash
# 监控容器日志
docker-compose logs -f wallet-app

# 检查安全事件
docker events --filter container=wallet-app
```

### 定期安全检查
```bash
# 扫描镜像漏洞
docker scan wallet-app:latest

# 检查容器配置
docker inspect wallet-app
```

## ⚠️ 重要安全提醒

1. **密码安全**
   - 必须修改所有默认密码
   - 使用强密码（包含大小写字母、数字、特殊字符）
   - 定期更换密码

2. **证书管理**
   - 生产环境建议使用HTTPS
   - 定期更新SSL证书
   - 妥善保管私钥

3. **数据备份**
   - 定期备份数据库
   - 备份文件加密存储
   - 测试恢复流程

4. **网络安全**
   - 使用VPN或内网访问
   - 配置适当的防火墙规则
   - 监控异常网络流量

5. **更新维护**
   - 定期更新基础镜像
   - 及时修复安全漏洞
   - 监控安全公告

## 🔍 安全检查清单

- [ ] 已修改所有默认密码
- [ ] 已配置防火墙规则
- [ ] 已启用日志监控
- [ ] 已测试备份恢复
- [ ] 已限制网络访问
- [ ] 已移除调试工具
- [ ] 已配置SSL证书（生产环境）
- [ ] 已设置监控告警

## 📞 安全事件响应

如发现安全问题：

1. **立即隔离**
   ```bash
   docker-compose down  # 停止所有服务
   ```

2. **备份证据**
   ```bash
   docker logs wallet-app > security-incident.log
   ```

3. **分析日志**
   ```bash
   grep -i "error\|fail\|attack" security-incident.log
   ```

4. **修复重部署**
   ```bash
   ./build.sh clean   # 清理环境
   ./build.sh build   # 重新构建部署
   ```

这些安全措施确保了容器在生产环境中的安全运行，有效防止了权限提升、恶意代码执行等安全风险。 