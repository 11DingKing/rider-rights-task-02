# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

朱家湖社区有一条维权事项已经结案，详情页也显示为结案，但再次调用状态流转接口时它又回到了办理中。请修复终态事项被重新打开的路径，正常结案和撤销仍要照常工作；测试代码不要跟着改。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-02
- 仓库地址：https://github.com/11DingKing/rider-rights-task-02.git
- parent SHA：8249a0de1b3004b44b3a7a3406d92eb79f3fec16

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-02.git bug-repro
cd bug-repro
git checkout --detach 8249a0de1b3004b44b3a7a3406d92eb79f3fec16
go test ./internal/domain -run "^TestTerminalCaseCannotReopen$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestTerminalCaseCannotReopen$" -count=1
--- FAIL: TestTerminalCaseCannotReopen (0.00s)
    terminal_transition_task_test.go:8: completed case was reopened
FAIL
FAIL	riderguard/internal/domain	0.040s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestTerminalCaseCannotReopen$" -count=1
--- FAIL: TestTerminalCaseCannotReopen (0.00s)
    terminal_transition_task_test.go:8: completed case was reopened
FAIL
FAIL	riderguard/internal/domain	0.001s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

已结案或已撤销的事项再次尝试进入办理中时必须返回非法状态转换错误并保持原终态；正常的登记、办理、结案和撤销流程仍须通过。定向测试、相关包测试和仓库全量回归必须通过，不得通过修改或跳过测试规避问题。
