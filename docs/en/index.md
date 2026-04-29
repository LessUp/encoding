---
layout: home

hero:
  name: CompressKit
  text: Compression algorithms you can trust
  tagline: Production-ready Huffman, Arithmetic, Range, and RLE implementations in C++, Go, and Rust. Verified across languages, documented for learning.
  image:
    src: /logo.svg
    alt: CompressKit Logo
  actions:
    - theme: brand
      text: Quick Start
      link: /en/guide/getting-started
    - theme: alt
      text: Compare Algorithms
      link: /en/guide/algorithms
    - theme: alt
      text: 中文
      link: /zh/

features:
  - icon: 🔄
    title: Cross-language verified
    details: Every algorithm is implemented in C++17, Go, and Rust with identical binary formats. Encode in one language, decode in another.
  - icon: 📚
    title: Learn by reading
    details: Clean, well-commented code designed for education. Each implementation fits in a single file you can actually read.
  - icon: ⚡
    title: Production-ready
    details: Security limits (4 GiB input, 1 GiB output), streaming APIs, comprehensive tests, and clear documentation.
  - icon: 🧪
    title: Test-driven quality
    details: 144 cross-language conformance tests ensure binary compatibility. Every release is verified before shipping.
---

<StatsBar />

## Why CompressKit?

| You need | We provide |
|----------|------------|
| Learn compression algorithms | Read implementations that fit in one file |
| Cross-language compatibility | Verified binary formats across C++/Go/Rust |
| Production use | Streaming APIs, security limits, error handling |
| Benchmark comparison | Run `make bench` and see real numbers |

## Algorithm Selection Guide

<AlgorithmGrid />

## Get started in 30 seconds

```bash
git clone https://github.com/LessUp/compress-kit.git
cd compress-kit
make build && make test
```

That's it. All 12 implementations (4 algorithms × 3 languages) built and tested.

## What makes this different

**Not another "compress everything" library.** CompressKit is a compression laboratory:

- **Transparent formats** — No opaque magic, every byte documented
- **Isomorphic implementations** — Same algorithm, same output, different languages
- **Educational focus** — Code you can read, tests you can trace
- **Verification-first** — Cross-language conformance is a test, not an afterthought

## Magic Numbers

Every algorithm has a 4-byte magic header for instant file identification:

| Algorithm | Magic | Description |
|-----------|-------|-------------|
| Huffman | `HFMN` | Prefix-code based compression |
| Arithmetic | `AENC` | Entropy-optimal encoding |
| Range Coder | `RCNC` | Fast integer arithmetic coding |
| RLE | `RLE\x00` | Run-length compression |

## Next Steps

| Goal | Page |
|------|------|
| Build and run locally | [Getting Started](/en/guide/getting-started) |
| Choose the right algorithm | [Algorithm Guide](/en/guide/algorithms) |
| Use as a library | [Streaming API](/en/api/streaming) |
| Verify compatibility | [Cross-Language Testing](/en/testing/cross-language) |
