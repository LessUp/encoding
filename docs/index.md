---
layout: home

hero:
  name: Encoding
  text: Compression Algorithms Collection
  tagline: Classic compression algorithms implemented in C++17, Go, and Rust for learning, comparison, and cross-language verification
  image:
    src: /logo.svg
    alt: Encoding Logo
  actions:
    - theme: brand
      text: Get Started →
      link: /en/guide/getting-started
    - theme: alt
      text: 中文文档
      link: /zh/guide/getting-started
    - theme: alt
      text: View on GitHub
      link: https://github.com/LessUp/encoding

features:
  - icon: 🌐
    title: Multi-Language Comparison | 多语言对比
    details: Each algorithm implemented in C++17, Go, and Rust. 每种算法都有 C++17、Go 和 Rust 实现。
  - icon: 📦
    title: Unified File Formats | 统一文件格式
    details: Cross-language compatible binary formats. 跨语言兼容的二进制格式。
  - icon: 📚
    title: Learning-Oriented | 面向学习
    details: Documentation focuses on algorithms and learning paths. 文档侧重于算法和学习路径。
  - icon: ✅
    title: Production-Ready | 生产级
    details: Complete CI/CD with automated testing. 完整的 CI/CD 与自动化测试。
---

## 🌍 Select Language | 选择语言

<div class="language-selector">

### [🇺🇸 English Documentation](/en/)

Complete documentation in English, including getting started guide, algorithm explanations, and project structure.

### [🇨🇳 中文文档](/zh/)

完整的中文文档，包括快速开始指南、算法详解和项目结构说明。

</div>

<style>
.language-selector {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
  margin: 2rem 0;
}

@media (max-width: 768px) {
  .language-selector {
    grid-template-columns: 1fr;
  }
}
</style>
