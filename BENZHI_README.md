# BENZHI_README

## 项目说明

- 项目：zhanglei10281852-gif/t63-qa-11
- 项目用途：Sanitation Operations is a Go backend for municipal vehicle dispatch and route operations. It models route planning, shift assignment, dispatch, inspection, incidents, fuel records, maintenance, audit events, outbox processing, and daily reconciliation.
- Go 工具链：`golang:1.22`
- 前端工具链：Node.js 20

## 标准构建、运行和测试命令

进入容器后执行：

```bash
# 编译
cd '/app' && GOTOOLCHAIN=local go build ./...
cd '/app/frontend' && npm install
cd '/app/frontend' && npm run build

# 启动
cd '/app' && GOTOOLCHAIN=local go run ./cmd/server
cd '/app/frontend' && npm run dev

# 测试
cd '/app' && GOTOOLCHAIN=local go test ./...
cd '/app/frontend' && npm test
```

## Docker 构建和进入容器

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh benzhi-task-11-amd64 linux/amd64
./build_benzhi_docker.sh benzhi-task-11-arm64 linux/arm64
docker run -it benzhi-task-11-amd64:latest
docker run -it --platform linux/arm64 benzhi-task-11-arm64:latest
```

## 题目验证命令

1. 预期退出码 1：`go test ./internal/worker -run TestRunnerCancelsInFlightHandler -count=1`

## Bug 复现

Bug 现象、触发步骤和完整错误信息见 `BUG_REPRO.md`。
