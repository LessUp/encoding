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
      text: Algorithms
      link: /en/guide/algorithms
    - theme: alt
      text: View on GitHub
      link: https://github.com/LessUp/encoding

features:
  - icon: 🌐
    title: Multi-Language Comparison
    details: Each algorithm is implemented in C++17, Go, and Rust, making it easy to compare code style, engineering practices, and performance characteristics across languages.
  - icon: 📦
    title: Unified File Formats
    details: All language implementations share identical binary formats, enabling direct cross-language encoding/decoding verification.
  - icon: 📚
    title: Learning-Oriented
    details: Documentation focuses on algorithm use cases, theoretical principles, and learning paths rather than just command lists.
  - icon: ✅
    title: Production-Ready Verification
    details: Complete CI/CD pipelines with automated builds, tests, and benchmarks to verify correctness and performance.
---

## 🎯 Project Overview

**Encoding** is an educational repository centered around classic compression algorithms. It provides working implementations alongside comprehensive documentation that explains algorithm backgrounds, applicable scenarios, and code organization.

### Who This Is For

| Audience | Use Case |
|----------|----------|
| 🎓 **Students & Learners** | Understand compression algorithms through multi-language comparison |
| 👨‍💻 **Software Engineers** | Compare C++, Go, and Rust implementation patterns for the same algorithm |
| 🔧 **Open Source Maintainers** | Verify cross-language format compatibility and benchmark performance |

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/LessUp/encoding.git
cd encoding

# Build all implementations
make build

# Run tests
make test

# Run benchmarks
make bench
```

## 📖 Documentation Structure

| Section | Description | Link |
|---------|-------------|------|
| **Getting Started** | Environment setup, build instructions, basic usage | [Read →](/en/guide/getting-started) |
| **Algorithms** | Algorithm explanations, comparisons, use cases | [Read →](/en/guide/algorithms) |
| **Project Structure** | Directory layout, CLI conventions, file formats | [Read →](/en/guide/project-structure) |
| **Changelog** | Version history and release notes | [View on GitHub](https://github.com/LessUp/encoding/blob/master/CHANGELOG.md) |

## 📊 Algorithm Overview

| Algorithm | Compression | Speed | Best For |
|-----------|-------------|-------|----------|
| **Huffman** | Medium | Fast | General purpose text/data |
| **Arithmetic** | High | Medium | Maximum compression needs |
| **Range Coder** | High | Fast | Balanced performance |
| **RLE** | Variable | Very Fast | Highly repetitive data |

## 🛠️ Tech Stack

- **C++17** - Zero-dependency, single-file implementations
- **Go 1.21+** - Module-based with library APIs
- **Rust 1.70+** - Cargo-based with library crates
- **Python 3.8+** - Benchmark and test scripts

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](https://github.com/LessUp/encoding/blob/master/CONTRIBUTING.md) for details.

## 📄 License

[MIT License](https://github.com/LessUp/encoding/blob/master/LICENSE) © 2025-2026 LessUp
