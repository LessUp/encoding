# 项目结构

## 总览

```
encoding/
├── huffman/              # Huffman 编码
│   ├── cpp/              #   C++ 单文件实现
│   ├── go/               #   Go 实现 (go.mod)
│   ├── rust/             #   Rust 实现
│   ├── benchmark/        #   跨语言基准测试
│   └── changelog/        #   变更记录
├── arithmetic/           # 算术编码
│   ├── cpp/              #   C++ 单文件实现
│   ├── go/               #   Go 实现
│   ├── rust/             #   Rust 实现
│   └── benchmark/        #   跨语言基准测试
├── range/                # 区间编码 (Range Coder)
│   ├── cpp/              #   C++ 单文件实现
│   ├── go/               #   Go 库 + CLI
│   ├── rust/             #   Rust 库 crate + CLI
│   └── benchmark/        #   跨语言基准测试
├── rle/                  # 游程编码 (RLE)
│   ├── cpp/              #   C++ 单文件实现
│   ├── go/               #   Go 实现
│   ├── rust/             #   Rust 实现
│   └── benchmark/        #   跨语言基准测试
├── scripts/              # 工具脚本 (run_all_bench.py)
├── tests/                # 测试数据生成 (gen_testdata.py)
├── docs/                 # VitePress 文档站
├── .github/workflows/    # CI + Pages 部署
├── Makefile              # 构建/测试/基准一体化入口
└── go.work               # Go workspace（多模块）
```

## 各语言实现约定

| 语言 | 版本 | 构建方式 | 特点 |
|------|------|---------|------|
| C++ | C++17 | `g++ -std=c++17 -O2` | 单文件实现，零依赖 |
| Go | 1.21+ | Go modules (`go.mod`) | Range Coder 提供 library API |
| Rust | 1.70+ | Cargo / rustc | Range Coder 提供 library crate |

## CLI 接口统一

所有实现遵循相同的 CLI 接口：

```bash
<algorithm>_<lang> encode <input> <output>
<algorithm>_<lang> decode <input> <output>
```

示例：`huffman_cpp`、`arithmetic_go`、`rangecoder_rust`、`rle_cpp`

## 文件格式兼容

同一算法的所有语言实现使用相同的二进制文件格式：

| 算法 | Magic | 扩展名 |
|------|-------|--------|
| Huffman | `HFMN` | `.huf` |
| Arithmetic | — | `.aenc` |
| Range Coder | — | `.rcnc` |
| RLE | — | `.rle` |

跨语言验证矩阵：

- C++ 编码 → Go 解码 ✓
- Go 编码 → Rust 解码 ✓
- Rust 编码 → C++ 解码 ✓

## CI/CD

| 工作流 | 文件 | 触发条件 |
|--------|------|---------|
| CI 测试 | `ci.yml` | 代码推送 / PR |
| Pages 部署 | `pages.yml` | `docs/` 变更 / 手动触发 |
