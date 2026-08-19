# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

关闭服务时 Runner 已收到取消信号，但正在执行的任务处理器一直不退出，优雅关闭最终超时。请先不要修改代码，定位取消信号在 worker 调用链中的断点并给出证据。

## 含 Bug 版本

- 仓库：zhanglei10281852-gif/t63-qa-11
- 仓库地址：https://github.com/zhanglei10281852-gif/t63-qa-11.git
- parent SHA：d75a70f6adc1890acf6de6ee5fb00f9d55c6ac5d

## 复现步骤

```bash
git clone -- https://github.com/zhanglei10281852-gif/t63-qa-11.git bug-repro
cd bug-repro
git checkout --detach d75a70f6adc1890acf6de6ee5fb00f9d55c6ac5d
go test ./internal/worker -run TestRunnerCancelsInFlightHandler -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/worker -run TestRunnerCancelsInFlightHandler -count=1
--- FAIL: TestRunnerCancelsInFlightHandler (1.81s)
    handler_cancellation_test.go:35: worker handler ignored cancellation
FAIL
FAIL	sanitation-operations/internal/worker	1.819s
FAIL

```

stderr：

```text
warning: internal/worker/handler_cancellation_test.go has type 100755, expected 100644
warning: internal/worker/worker_test.go has type 100755, expected 100644
warning: internal/worker/handler_cancellation_test.go has type 100755, expected 100644
warning: internal/worker/worker_test.go has type 100755, expected 100644

```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/worker -run TestRunnerCancelsInFlightHandler -count=1
--- FAIL: TestRunnerCancelsInFlightHandler (1.38s)
    handler_cancellation_test.go:35: worker handler ignored cancellation
FAIL
FAIL	sanitation-operations/internal/worker	1.623s
FAIL

```

stderr：

```text
warning: internal/worker/handler_cancellation_test.go has type 100755, expected 100644
warning: internal/worker/worker_test.go has type 100755, expected 100644
warning: internal/worker/handler_cancellation_test.go has type 100755, expected 100644
warning: internal/worker/worker_test.go has type 100755, expected 100644

```

## 通过条件

在触发条件下，定向测试 TestRunnerCancelsInFlightHandler 应通过，相关包、全量测试、竞态测试和构建检查均通过；回退 gold 唯一修复后定向测试重新失败。
