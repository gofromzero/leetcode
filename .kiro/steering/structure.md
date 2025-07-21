# 项目结构与组织规范

## 目录结构
```
├── algorithms/           # 算法实现目录
│   ├── go/              # Go语言实现
│   ├── rust/            # Rust语言实现  
│   └── cpp/             # C++语言实现（预留）
├── .github/             # GitHub配置和工作流
├── .kiro/               # Kiro AI助手配置
├── target/              # Rust构建输出目录
├── build/               # Go构建输出目录
├── Cargo.toml           # Rust工作区配置
├── go.mod               # Go模块配置
├── Makefile             # 构建脚本
└── README.md            # 项目文档
```

## 算法组织规范

### Go算法结构
- **目录命名**: 使用驼峰命名法，如`twosum`、`binarySearch`
- **文件命名**: 与目录名保持一致，如`twoSum.go`
- **包命名**: 与目录名相同
- **每个算法独立目录**: 便于管理和测试

### Rust算法结构
- **文件命名**: 使用下划线命名法，如`two_sum.rs`
- **模块组织**: 在`algorithms/rust/`目录下
- **二进制配置**: 在根目录`Cargo.toml`中配置

## 代码规范

### 文件头注释
Go文件应包含：
```go
// Copyright 2021 The GoLand Authors.
// Author: [作者名]
// Modified: [修改日期]
```

Rust文件应包含：
```rust
// LeetCode Problem: [题目名称]
// Author: [作者名]
// Language: Rust
```

### 命名约定
- **Go**: 使用驼峰命名法（camelCase）
- **Rust**: 使用下划线命名法（snake_case）
- **函数名**: 与LeetCode题目要求保持一致

### 测试规范
- **Go**: 测试文件命名为`*_test.go`
- **Rust**: 使用`#[cfg(test)]`模块进行测试
- **测试覆盖**: 确保核心算法逻辑有测试覆盖

## README维护
- **题目表格**: 按题号排序，包含题目链接、解决方案链接和难度
- **多语言链接**: 同一题目的不同语言实现都要在表格中体现
- **徽章状态**: 保持CI状态和覆盖率徽章更新

## 新增算法流程
1. 在对应语言目录下创建新文件夹
2. 实现算法解决方案
3. 添加必要的测试用例
4. 更新README.md中的题目表格
5. 确保CI通过所有检查