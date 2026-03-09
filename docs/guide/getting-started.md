# 快速开始

## 环境要求

| 工具 | 最低版本 | 用途 |
|------|---------|------|
| g++ / clang++ | 9+ / 10+ | C++17 实现编译 |
| Go | 1.21+ | Go 实现 |
| Rust (cargo) | 1.70+ | Rust 实现 |
| Python | 3.8+ | 基准测试脚本 |

## 克隆与构建

```bash
git clone https://github.com/LessUp/encoding.git
cd encoding
```

### 一键构建所有实现

```bash
make build
```

### 单独构建某个算法

```bash
make build-huffman      # Huffman (C++, Go, Rust)
make build-arithmetic   # 算术编码
make build-range        # 区间编码
make build-rle          # 游程编码
```

### 手动编译示例

::: code-group
```bash [C++]
cd huffman/cpp
g++ -std=c++17 -O2 main.cpp -o huffman_cpp
./huffman_cpp encode input.bin output.huf
./huffman_cpp decode output.huf restored.bin
```

```bash [Go]
cd huffman/go
go build -o huffman_go .
./huffman_go encode input.bin output.huf
```

```bash [Rust]
cd huffman/rust
rustc -O main.rs -o huffman_rust
./huffman_rust encode input.bin output.huf
```
:::

## 跨语言验证

所有实现使用相同文件格式，可以交叉验证：

```bash
# C++ 编码，Go 解码
./huffman_cpp encode input.bin encoded.huf
./huffman_go decode encoded.huf decoded.bin
diff input.bin decoded.bin  # 无输出表示一致
```

支持任意方向的组合：C++ ↔ Go ↔ Rust。

## 运行测试

```bash
make test          # 运行所有 Go + Rust 单元测试
```

## 运行基准测试

```bash
make bench         # 自动生成测试数据 + 运行所有基准
```

报告输出到 `reports/` 目录。

## Makefile 命令速查

| 命令 | 说明 |
|------|------|
| `make build` | 构建所有语言的所有算法实现 |
| `make test` | 运行所有 Go 和 Rust 单元测试 |
| `make bench` | 生成测试数据并运行跨语言基准测试 |
| `make test-data` | 仅生成测试数据 |
| `make clean` | 清理所有构建产物和报告 |
