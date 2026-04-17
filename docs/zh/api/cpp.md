# C++ 实现参考

所有 C++ 实现都是单文件、零依赖的 C++17 程序。

## 编译

```bash
g++ -std=c++17 -O2 -Wall -Wextra -o <binary> main.cpp
```

### 推荐编译选项

| 选项 | 用途 |
|------|------|
| `-std=c++17` | 启用 C++17 特性 |
| `-O2` | 优化级别 |
| `-Wall -Wextra` | 警告 |
| `-fsanitize=address` | 地址消毒器（调试版本） |
| `-fsanitize=undefined` | 未定义行为消毒器（调试版本） |

## Huffman 编码

```bash
./huffman_cpp encode input.bin output.huf
./huffman_cpp decode output.huf decoded.bin
```

### 内部结构

- `BitWriter` / `BitReader` — 比特级 I/O 类
- `Node` — Huffman 树节点（自定义删除器）
- `compress_file()` / `decompress_file()` — 主编/解码函数

### 文件格式

| 偏移 | 大小 | 字段 |
|------|------|------|
| 0 | 4B | 魔数: `HFMN` |
| 4 | 4B | 频率表大小（始终 257） |
| 8 | 1028B | 频率表（257 × uint32 LE） |
| 1036+ | 可变 | 编码比特流 |

---

## 算术编码 / 区间编码 / RLE

遵循相同的单文件、零依赖模式。CLI 接口统一为：

```bash
<algorithm>_cpp encode <input> <output>
<algorithm>_cpp decode <input> <output>
```

### 通用模式

| 模式 | 描述 |
|------|------|
| 单文件 | 每个算法在一个 `main.cpp` 中 |
| 零依赖 | 仅使用标准库 |
| 错误处理 | `fprintf(stderr, ...)` + `exit(1)` |
| 内存管理 | `std::unique_ptr` + 自定义删除器 |
