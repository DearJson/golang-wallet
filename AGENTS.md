# Repository Guidelines

## Project Structure & Module Organization
核心入口在 `main.go`，按领域划分的业务模块位于 `app/`（如 `app/system`、`app/web`），共享逻辑集中在 `app/common`。链路交互和合约定义分别在 `rpc/`、`contract/` 与 `abi/` 中。统一配置放在 `config/`，中间件与路由位于 `middleware/`、`router/`。静态资源与模板分别在 `public/`、`template/`，Docker 与部署脚本集中在 `docker/` 和 `build.sh`。

## Build, Test, and Development Commands
使用 `go mod tidy` 维护依赖；日常构建执行 `go build ./...`。本地运行可用 `go run main.go`，需加载 `config/` 下的环境配置。回归测试执行 `go test ./...`，CI 前建议附加 `go vet ./...`。容器化流程通过 `./build.sh` 或 `docker-compose --env-file docker.env up -d wallet-app` 启动主要服务。

## Coding Style & Naming Conventions
遵循 Go 官方格式，提交前执行 `gofmt -w` 和（可选）`goimports`。包名保持小写单词，文件命名使用下划线连接领域与用途（例如 `wallet_service.go`）。导出结构使用驼峰命名并提供简洁注释，错误变量以 `Err` 前缀命名。

## Testing Guidelines
现有单测位于 `hdwallet/core_test.go`，新增测试应置于同一模块目录并以 `_test.go` 结尾。优先采用表驱动用例覆盖边界场景。确保 `go test ./...` 全绿，再补充 `-run` 过滤器验证关键路径。若引入外部依赖，使用接口或模拟以保持测试可重复。

## Commit & Pull Request Guidelines
Git 历史以简短英文动词开头（如 `update .gitignore`、`fix`），合并前请确保信息清晰、上下文完整。推荐格式为 `scope: action`，必要时附加问题编号。PR 应包含：变更摘要、测试结果、受影响模块列表以及与接口或配置相关的截图或日志。

## Security & Configuration Tips
敏感配置集中在 `docker.env` 与 `config/`，禁止提交明文秘钥。首次部署请运行 `./build.sh` 以自动生成安全随机字符串，并核对 `docker` 目录中的安全指引。若更新合约或 ABI，保持 `contract/`、`abi/` 与对应 Go 绑定同步，避免节点与链上接口失配。
