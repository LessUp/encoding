# 项目结构

## 总览

```
encoding/
├── huffman/           # Huffman 编码
│   ├── cpp/          # C++ 实现
│   ├── go/           # Go 实现
│   ├── rust/         # Rust 实现
│   └── benchmark/    # 跨语言基准测试
├── arithmetic/        # 算术编码
│   ├── cpp/          # C++ 实现
│   ├── go/           # Go 实现
│   ├── rust/         # Rust 实现
│   └── benchmark/    # 跨语言基准测试
├── range/            # 区间编码
│   ├── cpp/          # C++ 实现
│   ├── go/           # Go（库 + CLI）
│   ├── rust/         # Rust（库 + CLI）
│   └── benchmark/    # 跨语言基准测试
├── rle/              # 游程编码
│   ├── cpp/          # C++ 实现
│   ├── go/           # Go 实现
│   ├── rust/         # Rust 实现
│   └── benchmark/    # 跨语言基准测试
├── scripts/          # 工具脚本
└── tests/            # 测试数据生成
```

## 各语言实现约定

### C++

- 标准：C++17
- 编译：`g++ -std=c++17 -O2 main.cpp -o <name>_cpp`
- 单文件实现，便于阅读

### Go

- 版本：Go 1.21+
- 使用 Go modules（`go.mod`）
- Range Coder 额外提供 library API

### Rust

- 版本：Rust 1.70+
- 使用 Cargo（`Cargo.toml`）
- Range Coder 额外提供 library crate

## CLI 接口统一

所有实现遵循相同的 CLI 接口：

```bash
<algorithm>_<lang> encode <input> <output>
<algorithm>_<lang> decode <input> <output>
```

## 文件格式兼容

同一算法的所有语言实现使用相同的二进制文件格式，确保：

- C++ 编码 → Go 解码 ✓
- Go 编码 → Rust 解码 ✓
- Rust 编码 → C++ 解码 ✓
