---
layout: home

hero:
  name: CompressKit
  text: 经典压缩算法，三语言实现
  tagline: Huffman、算术编码、Range Coder 和 RLE 的 C++、Go、Rust 实现。跨语言验证，文档齐全，适合学习与生产。
  image:
    src: /logo.svg
    alt: CompressKit Logo
  actions:
    - theme: brand
      text: 快速开始
      link: /zh/guide/getting-started
    - theme: alt
      text: 算法对比
      link: /zh/guide/algorithms
    - theme: alt
      text: English
      link: /en/

features:
  - icon: 🔄
    title: 跨语言验证
    details: 每个算法都有 C++17、Go、Rust 三种实现，二进制格式完全一致。一种语言编码，另一种语言解码。
  - icon: 📚
    title: 可读性优先
    details: 代码整洁、注释完善，专为教学设计。每个实现都控制在单文件内，真正可以读完。
  - icon: ⚡
    title: 生产可用
    details: 安全限制（4 GiB 输入、1 GiB 输出）、流式 API、完整测试、清晰文档。
  - icon: 🧪
    title: 测试驱动
    details: 144 项跨语言一致性测试确保二进制兼容。每次发布前都经过完整验证。
---

<StatsBar />

## 为什么选择 CompressKit？

| 你的需求 | 我们的方案 |
|----------|------------|
| 学习压缩算法 | 单文件实现，代码可读 |
| 跨语言兼容 | C++/Go/Rust 二进制格式一致 |
| 生产使用 | 流式 API、安全限制、错误处理 |
| 性能对比 | 运行 `make bench` 获取真实数据 |

## 算法选择指南

<AlgorithmGrid />

## 30 秒快速上手

```bash
git clone https://github.com/LessUp/compress-kit.git
cd compress-kit
make build && make test
```

搞定。所有 12 个实现（4 种算法 × 3 种语言）已构建并测试通过。

## 独特之处

**不是另一个"万能压缩"库。** CompressKit 是一个压缩算法实验室：

- **透明格式** — 没有黑盒魔法，每个字节都有文档
- **同构实现** — 同一算法、相同输出、不同语言
- **教学导向** — 代码可读，测试可追踪
- **验证优先** — 跨语言一致性是测试，不是附加功能

## 魔数标识

每个算法都有 4 字节魔数头部，可即时识别文件类型：

| 算法 | 魔数 | 说明 |
|------|------|------|
| Huffman | `HFMN` | 前缀码压缩 |
| 算术编码 | `AENC` | 熵最优编码 |
| Range Coder | `RCNC` | 快速整数算术编码 |
| RLE | `RLE\x00` | 行程编码 |

## 下一步

| 目标 | 页面 |
|------|------|
| 本地构建运行 | [快速开始](/zh/guide/getting-started) |
| 选择合适的算法 | [算法指南](/zh/guide/algorithms) |
| 作为库使用 | [流式 API](/zh/api/streaming) |
| 验证兼容性 | [跨语言测试](/zh/testing/cross-language) |
