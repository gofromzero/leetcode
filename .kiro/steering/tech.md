# 技术栈与构建系统

## 编程语言
- **Go**: 主要实现语言，版本要求 Go 1.22+
- **Rust**: 次要实现语言，使用2021版本
- **C++**: 预留支持，目前为空

## 构建系统

### Go项目构建
- **模块管理**: 使用Go Modules (`go.mod`)
- **工作区**: 项目根目录作为Go模块
- **依赖管理**: `go mod download`

### Rust项目构建
- **包管理**: 使用Cargo工具链
- **工作区**: 配置了Rust workspace，成员包括 `algorithms/rust`
- **二进制目标**: 配置了`two_sum`二进制文件

## 常用命令

### 开发命令
```bash
# 安装依赖
make dep

# 代码检查
make lint
make vet

# 运行测试
make test

# 测试覆盖率
make test-coverage

# 构建项目
make build

# 清理构建文件
make clean
```

### Go特定命令
```bash
# 下载依赖
go mod download

# 运行测试
go test ./...

# 代码格式化
go fmt ./...

# 静态分析
go vet ./...
```

### Rust特定命令
```bash
# 构建Rust代码
cargo build

# 运行测试
cargo test

# 运行特定二进制
cargo run --bin two_sum
```

## 持续集成
- **GitHub Actions**: 配置了自动化测试和构建
- **代码覆盖率**: 使用codecov进行覆盖率统计
- **质量检查**: 集成了代码质量检查工具

## 开发工具
- **Makefile**: 提供统一的构建和测试命令
- **Git**: 版本控制，配置了.gitignore
- **IDE支持**: 包含.idea配置文件夹