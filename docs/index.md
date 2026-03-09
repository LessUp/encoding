---
layout: home

hero:
  name: Encoding
  text: 编码算法集合
  tagline: 用 C++、Go、Rust 多语言实现经典压缩编码算法，学习与对比
  actions:
    - theme: brand
      text: 快速开始
      link: /guide/getting-started
    - theme: alt
      text: GitHub
      link: https://github.com/LessUp/encoding

features:
  - title: Huffman 编码
    details: 基于前缀码的无损压缩，通用文本压缩，中等压缩率，速度快
    icon: 🌳
  - title: 算术编码
    details: 区间逐步细分表示消息概率，压缩效率最接近信息熵上界
    icon: 📐
  - title: 区间编码 (Range Coder)
    details: 等价于算术编码但实践中更高效，平衡压缩率与速度
    icon: 📏
  - title: 游程编码 (RLE)
    details: 适用于大量连续重复字节的数据，速度极快
    icon: 🔁
  - title: 多语言实现
    details: 每种算法均提供 C++17、Go、Rust 三种实现，便于对比学习
    icon: 🌐
  - title: 跨语言兼容
    details: 所有实现使用相同文件格式，支持交叉编码/解码验证
    icon: 🔄
---
