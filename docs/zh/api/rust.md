# Rust Crate API 参考

另请参阅: [Streaming API](/zh/api/streaming)

每种 Rust 实现都提供可复用的库 crate，现在四种算法也统一提供 shared streaming adapter 与 `compresskit-codec` buffer helper。

## Huffman (`algorithms/huffman/rust`)

### 添加依赖

```toml
[dependencies]
huffman = { path = "../algorithms/huffman/rust" }
```

### API

#### `huffman_encode_file(input: &str, output: &str) -> io::Result<()>`

使用 Huffman 编码压缩文件。

```rust
use huffman;

fn main() -> std::io::Result<()> {
    huffman::huffman_encode_file("input.bin", "output.huf")?;
    huffman::huffman_decode_file("output.huf", "decoded.bin")?;
    Ok(())
}
```

### 内部函数

- `compress_file(input: &str) -> io::Result<Vec<u8>>` — 返回编码后的字节
- `decompress_file(input: &str) -> io::Result<Vec<u8>>` — 返回解码后的字节

---

## Arithmetic (`algorithms/arithmetic/rust`)

### API

```rust
use arithmetic;

fn main() -> std::io::Result<()> {
    arithmetic::arithmetic_encode_file("input.bin", "output.aenc")?;
    arithmetic::arithmetic_decode_file("output.aenc", "decoded.bin")?;
    Ok(())
}
```

---

## Range Coder (`algorithms/range/rust`)

### 添加依赖

```toml
[dependencies]
rangecoder = { path = "../algorithms/range/rust" }
```

### API

#### `encode(data: &[u8]) -> Result<Vec<u8>, RangeError>`

使用 Range Coder 编码数据。

#### `decode(data: &[u8]) -> Result<Vec<u8>, RangeError>`

解码 Range Coder 编码的数据。

```rust
use rangecoder;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let encoded = rangecoder::encode(&input_data)?;
    let decoded = rangecoder::decode(&encoded)?;
    Ok(())
}
```

### 错误类型

```rust
pub enum RangeError {
    InvalidHeader(String),
    CorruptedData(String),
    IoError(std::io::Error),
}
```

### CLI 二进制文件

```bash
cargo build --bin rangecoder --release
./rangecoder encode input.bin output.rcnc
./rangecoder decode output.rcnc decoded.bin
```

---

## RLE (`algorithms/rle/rust`)

### API

```rust
use rle;

fn main() -> std::io::Result<()> {
    rle::rle_encode_file("input.bin", "output.rle")?;
    rle::rle_decode_file("output.rle", "decoded.bin")?;
    Ok(())
}
```

---

## 通用模式

所有 Rust 实现都遵循以下约定：

| 模式 | 说明 |
|------|------|
| `*_encode_file` | 文件编码，返回 `io::Result<()>` |
| `*_decode_file` | 文件解码，返回 `io::Result<()>` |
| 错误传播 | 全程使用 `?` 运算符 |
| 缓冲 I/O | 所有文件操作使用 `BufReader`/`BufWriter` |

### 错误处理

- `io::Error` — 文件 I/O 失败
- `Box<dyn Error>` — 算法特定错误（Range Coder）
- 所有错误向上传播到 `main()` 并输出到 stderr
