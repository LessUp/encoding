# 算法详解

## 算法对比总览

| 算法 | 适用场景 | 压缩率 | 速度 | 时间复杂度 | 空间复杂度 |
|------|----------|--------|------|-----------|-----------|
| **Huffman** | 通用、文本 | 中等 | 快 | O(n log n) | O(σ) |
| **Arithmetic** | 追求最大压缩 | 高 | 中等 | O(n) | O(σ) |
| **Range Coder** | 平衡性能 | 高 | 快 | O(n) | O(σ) |
| **RLE** | 重复数据 | 可变 | 非常快 | O(n) | O(1) |

> σ = 字母表大小（字节级编码为 256），n = 输入长度

---

## Huffman 编码

基于前缀码的无损压缩算法。实现中先扫描输入统计频率，构建 Huffman 树，再按位写入编码结果。

### 原理

1. 统计输入中每个字节的出现频率
2. 构建 Huffman 二叉树（频率低的深度更深）
3. 生成前缀码表（无二义性）
4. 按码表编码输入，写入 bit stream

### 文件格式

| 字段 | 大小 | 说明 |
|------|------|------|
| Magic | 4 bytes | `HFMN` |
| 频率表 | 257 × 4 bytes | little-endian |
| 编码数据 | 变长 | bit stream |

### 压缩效率

- **理论下界**：每个符号平均编码长度 ≥ 信息熵 H
- **Huffman 上界**：H ≤ L < H + 1（每符号最多多 1 bit）
- 对频率分布越不均匀的数据效果越好

---

## 算术编码 (Arithmetic Coding)

使用区间逐步细分表示整段消息的概率，压缩效率更接近信息熵上界。

### 原理

1. 初始区间为 `[0, 1)`
2. 根据符号概率逐步细分区间
3. 最终用一个落在目标区间内的数表示整段消息
4. 输出比特数接近 `-log₂(P(message))`

### 与 Huffman 的对比

| | Huffman | Arithmetic |
|---|---------|-----------|
| 编码单位 | 每符号至少 1 bit | 可输出分数 bit |
| 理论效率 | H ≤ L < H + 1 | L ≈ H + ε |
| 实现难度 | 较低 | 较高（精度管理） |
| 适用场景 | 通用 | 高压缩率需求 |

### 特点

- 理论上压缩率最优（接近熵）
- 编解码速度较 Huffman 慢
- 实现复杂度较高（需处理精度问题）

---

## 区间编码 (Range Coder)

一种等价于算术编码的实现方式，但常在实践中更高效。核心区别在于使用整数区间运算替代浮点，避免专利问题的同时获得更好的实际性能。

### 与算术编码的区别

- 算术编码按 bit 输出，Range Coder 按 byte 输出 → I/O 更高效
- 压缩率与算术编码几乎相同
- 无专利限制，工程实践首选

### API

**Go**:
```go
encoded, err := rangecoder.Encode(data)
decoded, err := rangecoder.Decode(encoded)
```

**Rust**:
```rust
let encoded = rangecoder::encode(input)?;
let decoded = rangecoder::decode(&encoded)?;
```

### CLI

```bash
# C++
./rangecoder_cpp encode input.bin output.rcnc
./rangecoder_cpp decode output.rcnc restored.bin

# Go
./rangecoder_go encode input.bin output.rcnc

# Rust
cargo run --bin rangecoder -- encode input.bin output.rcnc
```

---

## 游程编码 (RLE)

适用于包含大量相同字节连续重复的数据。是最简单也是最古老的压缩算法之一。

### 文件格式

重复的 `(count, value)` 对：

| 字段 | 大小 | 说明 |
|------|------|------|
| count | 4 bytes | little-endian, unsigned |
| value | 1 byte | 原始字节 |

### 特点

- 实现最简单，编解码速度极快
- 对重复数据效果极佳（如位图、日志中的重复行）
- 对随机数据可能"膨胀"（输出比输入大，最坏膨胀 5×）
- 常作为其他压缩算法的预处理步骤（如 BWT + MTF + RLE + 算术编码）
