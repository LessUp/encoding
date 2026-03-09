---
layout: home

hero:
  name: Encoding
  text: 编码算法集合
  tagline: 用 C++17、Go、Rust 三种语言实现 Huffman、算术编码、区间编码、RLE 四大经典压缩算法 — 学习、对比、跨语言验证
  actions:
    - theme: brand
      text: 快速开始
      link: /guide/getting-started
    - theme: alt
      text: 算法详解
      link: /guide/algorithms
    - theme: alt
      text: GitHub
      link: https://github.com/LessUp/encoding

features:
  - title: Huffman 编码
    details: 基于前缀码的经典无损压缩。扫描频率 → 构建 Huffman 树 → 按位输出，通用文本场景首选
    icon: 🌳
    link: /guide/algorithms#huffman-编码
  - title: 算术编码
    details: 区间逐步细分表示消息概率，压缩效率最接近信息熵理论上界，追求极致压缩率时使用
    icon: 📐
    link: /guide/algorithms#算术编码-arithmetic-coding
  - title: 区间编码 (Range Coder)
    details: 等价于算术编码但实践中更高效，在压缩率与编解码速度之间取得最佳平衡
    icon: 📏
    link: /guide/algorithms#区间编码-range-coder
  - title: 游程编码 (RLE)
    details: 最简单的压缩算法，适用于大量连续重复字节的数据，编解码速度极快
    icon: 🔁
    link: /guide/algorithms#游程编码-rle
  - title: 三语言实现
    details: 每种算法均提供 C++17、Go 1.21+、Rust 1.70+ 三种实现，统一 CLI 接口，便于横向对比
    icon: 🌐
    link: /guide/project-structure
  - title: 跨语言兼容验证
    details: 所有实现共用相同二进制文件格式 — C++ 编码 → Go 解码 → Rust 再编码，保证正确性
    icon: 🔄
    link: /guide/getting-started#跨语言验证
---
