# 快速开始

## 环境要求

- C++ compiler (g++ 9+ 或 clang++ 10+)
- Go 1.21+
- Rust 1.70+
- Python 3.8+ (用于基准测试)

## 克隆与构建

```bash
git clone https://github.com/LessUp/encoding.git
cd encoding
```

### Huffman 编码示例

```bash
cd huffman/cpp
g++ -std=c++17 -O2 main.cpp -o huffman_cpp

# 编码
./huffman_cpp encode input.bin output.huf

# 解码
./huffman_cpp decode output.huf restored.bin
```

### Range Coder 示例

```bash
# C++
cd range/cpp && g++ -std=c++17 -O2 main.cpp -o rangecoder_cpp
./rangecoder_cpp encode input.bin output.rcnc

# Go
cd range/go && go build -o rangecoder_go ./cmd/rangecoder
./rangecoder_go encode input.bin output.rcnc

# Rust
cd range/rust && cargo run --bin rangecoder -- encode input.bin output.rcnc
```

## 跨语言验证

所有实现使用相同文件格式，可以交叉验证：

```bash
# C++ 编码，Go 解码
./huffman_cpp encode input.bin encoded.huf
./huffman_go decode encoded.huf decoded.bin
diff input.bin decoded.bin  # 无输出表示一致
```

## 运行基准测试

```bash
python3 scripts/run_all_bench.py
```

报告输出到 `reports/` 目录。

## 便捷命令 (Makefile)

```bash
make help          # 显示所有可用命令
make build-all     # 构建所有实现
make test-all      # 运行所有测试
make bench-all     # 运行所有基准测试
make clean         # 清理构建产物
```
