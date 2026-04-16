---
layout: home

hero:
  name: Encoding
  text: 压缩算法集
  tagline: 使用 C++17、Go 和 Rust 实现的经典压缩算法，用于学习、对比和跨语言验证
  image:
    src: /logo.svg
    alt: Encoding Logo
  actions:
    - theme: brand
      text: 快速开始 →
      link: /zh/guide/getting-started
    - theme: alt
      text: 算法详解
      link: /zh/guide/algorithms
    - theme: alt
      text: 查看 GitHub
      link: https://github.com/LessUp/encoding

features:
  - icon: 🌐
    title: 多语言对比
    details: 每种算法都有 C++17、Go 和 Rust 实现，便于对比代码风格、工程实践和性能特征。
  - icon: 📦
    title: 统一文件格式
    details: 所有语言实现共享相同的二进制格式，可直接进行跨语言编码/解码验证。
  - icon: 📚
    title: 面向学习
    details: 文档侧重于算法用例、理论原理和学习路径，而不仅仅是命令列表。
  - icon: ✅
    title: 生产级验证
    details: 完整的 CI/CD 流水线，自动化构建、测试和基准测试，确保正确性和性能。
---

## 🎯 项目简介

**Encoding** 是一个围绕经典压缩算法的教育性仓库。它提供可运行的实现以及全面的文档，解释算法背景、适用场景和代码组织。

### 适用人群

| 用户类型 | 使用场景 |
|----------|----------|
| 🎓 **学生和学习者** | 通过多语言对比理解压缩算法 |
| 👨‍💻 **软件工程师** | 对比相同算法在 C++、Go 和 Rust 中的实现模式 |
| 🔧 **开源维护者** | 验证跨语言格式兼容性和基准性能 |

## 🚀 快速开始

```bash
# 克隆仓库
git clone https://github.com/LessUp/encoding.git
cd encoding

# 构建所有实现
make build

# 运行测试
make test

# 运行基准测试
make bench
```

## 📖 文档结构

| 章节 | 描述 | 链接 |
|------|------|------|
| **快速开始** | 环境配置、构建说明、基本用法 | [阅读 →](/zh/guide/getting-started) |
| **算法详解** | 算法说明、对比、用例 | [阅读 →](/zh/guide/algorithms) |
| **项目结构** | 目录结构、CLI 规范、文件格式 | [阅读 →](/zh/guide/project-structure) |
| **更新日志** | 版本历史和发布说明 | [GitHub 查看](https://github.com/LessUp/encoding/blob/master/CHANGELOG.md) |

## 📊 算法概览

| 算法 | 压缩率 | 速度 | 适用场景 |
|------|--------|------|----------|
| **Huffman** | 中等 | 快 | 通用文本/数据 |
| **算术编码** | 高 | 中等 | 最大压缩需求 |
| **区间编码** | 高 | 快 | 平衡性能 |
| **RLE** | 可变 | 极快 | 高度重复数据 |

## 🛠️ 技术栈

- **C++17** - 无依赖单文件实现
- **Go 1.21+** - 基于模块，提供库 API
- **Rust 1.70+** - 基于 Cargo，提供库包
- **Python 3.8+** - 基准测试和测试脚本

## 🤝 参与贡献

我们欢迎贡献！详情请查看我们的 [贡献指南](https://github.com/LessUp/encoding/blob/master/CONTRIBUTING.md)。

## 📄 许可证

[MIT 许可证](https://github.com/LessUp/encoding/blob/master/LICENSE) © 2025-2026 LessUp
