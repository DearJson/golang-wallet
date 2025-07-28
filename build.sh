#!/bin/bash

# ===========================================
# 钱包管理系统 Docker 构建部署脚本
# ===========================================

set -e  # 遇到错误立即退出

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查Docker和Docker Compose
check_prerequisites() {
    log_info "检查系统依赖..."
    
    if ! command -v docker &> /dev/null; then
        log_error "Docker 未安装，请先安装 Docker"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        log_error "Docker Compose 未安装，请先安装 Docker Compose"
        exit 1
    fi
    
    log_success "系统依赖检查通过"
}

# 清理旧的构建
cleanup_old_builds() {
    log_info "清理旧的Docker构建..."
    
    # 停止并删除旧容器
    docker-compose down --remove-orphans 2>/dev/null || true
    
    # 删除旧镜像
    docker images | grep "golang-wallet" | awk '{print $3}' | xargs -r docker rmi -f 2>/dev/null || true
    
    # 清理构建缓存
    docker builder prune -f 2>/dev/null || true
    
    log_success "清理完成"
}

# 构建应用
build_application() {
    log_info "开始构建应用镜像..."
    
    # 确保必要的目录存在
    mkdir -p data logs
    
    # 设置目录权限
    chmod 755 data logs
    
    # 构建镜像
    docker-compose build --no-cache wallet-app
    
    if [ $? -eq 0 ]; then
        log_success "应用镜像构建成功"
    else
        log_error "应用镜像构建失败"
        exit 1
    fi
}

# 生成安全密码
generate_secure_passwords() {
    log_info "生成安全配置..."
    
    if [ ! -f "docker.env" ]; then
        log_error "docker.env 文件不存在"
        exit 1
    fi
    
    # 检查是否需要生成新密码
    if grep -q "your_secure.*password_here" docker.env; then
        log_warning "检测到默认密码，强烈建议修改 docker.env 中的密码配置"
        read -p "是否自动生成随机密码？(y/n): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            # 生成随机密码
            MYSQL_ROOT_PWD=$(openssl rand -base64 32)
            MYSQL_PWD=$(openssl rand -base64 32)
            REDIS_PWD=$(openssl rand -base64 32)
            RABBITMQ_PWD=$(openssl rand -base64 32)
            JWT_SECRET=$(openssl rand -base64 64)
            ENCRYPT_KEY=$(openssl rand -base64 32 | head -c 32)
            
            # 更新配置文件
            sed -i.bak "s/your_secure_root_password_here/${MYSQL_ROOT_PWD}/g" docker.env
            sed -i.bak "s/your_secure_wallet_password_here/${MYSQL_PWD}/g" docker.env
            sed -i.bak "s/your_secure_redis_password_here/${REDIS_PWD}/g" docker.env
            sed -i.bak "s/your_secure_rabbitmq_password_here/${RABBITMQ_PWD}/g" docker.env
            sed -i.bak "s/your_very_long_and_secure_jwt_secret_key_here/${JWT_SECRET}/g" docker.env
            sed -i.bak "s/your_32_character_encryption_key_here/${ENCRYPT_KEY}/g" docker.env
            
            log_success "随机密码已生成并保存到 docker.env"
        fi
    fi
}

# 启动服务
start_services() {
    log_info "启动服务..."
    
    # 使用环境变量文件启动
    docker-compose --env-file docker.env up -d
    
    if [ $? -eq 0 ]; then
        log_success "服务启动成功"
    else
        log_error "服务启动失败"
        exit 1
    fi
}

# 等待服务就绪
wait_for_services() {
    log_info "等待服务就绪..."
    
    # 等待MySQL就绪
    log_info "等待MySQL服务..."
    until docker-compose exec mysql mysqladmin ping -h"localhost" --silent; do
        sleep 2
    done
    
    # 等待Redis就绪
    log_info "等待Redis服务..."
    until docker-compose exec redis redis-cli ping | grep PONG; do
        sleep 2
    done
    
    # 等待RabbitMQ就绪
    log_info "等待RabbitMQ服务..."
    until docker-compose exec rabbitmq rabbitmqctl status; do
        sleep 2
    done
    
    # 等待应用就绪
    log_info "等待应用服务..."
    sleep 10
    until curl -f http://localhost:9090/ &>/dev/null; do
        sleep 5
    done
    
    log_success "所有服务已就绪"
}

# 显示服务状态
show_status() {
    log_info "服务状态："
    docker-compose ps
    
    echo
    log_info "服务访问地址："
    echo "  • 钱包管理系统: http://localhost:9090"
    echo "  • RabbitMQ管理界面: http://localhost:15672"
    echo "  • MySQL端口: localhost:3307"
    echo "  • Redis端口: localhost:6379"
    
    echo
    log_info "安全提醒："
    echo "  • 请及时修改默认密码"
    echo "  • 建议开启防火墙，只允许必要的端口访问"
    echo "  • 定期备份数据"
    echo "  • 监控日志文件"
}

# 显示日志
show_logs() {
    log_info "显示服务日志（按Ctrl+C退出）:"
    docker-compose logs -f
}

# 主函数
main() {
    echo "=========================================="
    echo "     钱包管理系统 Docker 部署脚本"
    echo "=========================================="
    
    case "${1:-build}" in
        "build")
            check_prerequisites
            cleanup_old_builds
            generate_secure_passwords
            build_application
            start_services
            wait_for_services
            show_status
            ;;
        "start")
            docker-compose --env-file docker.env up -d
            show_status
            ;;
        "stop")
            docker-compose down
            log_success "服务已停止"
            ;;
        "restart")
            docker-compose down
            docker-compose --env-file docker.env up -d
            show_status
            ;;
        "logs")
            show_logs
            ;;
        "status")
            show_status
            ;;
        "clean")
            cleanup_old_builds
            log_success "清理完成"
            ;;
        *)
            echo "使用方法: $0 {build|start|stop|restart|logs|status|clean}"
            echo
            echo "命令说明："
            echo "  build   - 构建并启动所有服务（默认）"
            echo "  start   - 启动服务"
            echo "  stop    - 停止服务"
            echo "  restart - 重启服务"
            echo "  logs    - 查看日志"
            echo "  status  - 查看状态"
            echo "  clean   - 清理旧构建"
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@" 