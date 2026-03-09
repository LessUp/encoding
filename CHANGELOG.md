# Changelog | 变更日志

All notable changes to this project will be documented in this file.

本文件记录本项目的所有重要变更。

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/)，
本项目遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

## [Unreleased]

### Added | 新增
- VitePress SEO: Open Graph / Twitter Card meta tags, sitemap, cleanUrls
- Docs landing page: feature cards with anchor links, richer hero tagline
- Docs algorithms page: complexity analysis table, Huffman/Arithmetic comparison, Range Coder vs Arithmetic differences
- Docs getting-started: code-group tabs, Makefile command table, structured environment requirements
- Docs project-structure: CI/CD table, file format table, enhanced directory tree
- Site favicon (logo.svg)
- Arithmetic coding Go implementation with full test suite
- Arithmetic coding Rust implementation with full test suite
- Range coder Go CLI (`range/go/cmd/main.go`) for cross-language testing
- Range coder Rust CLI (`range/rust/src/bin/rangecoder.rs`) for cross-language testing
- Cross-language benchmark script for Arithmetic coding (`arithmetic/benchmark/bench.py`)
- Cross-language benchmark script for Range coder (`range/benchmark/bench.py`)
- CI cross-language correctness tests for Arithmetic and Range coder (all 3 languages)

### Changed | 变更
- Pages workflow: sparse-checkout (skip algorithm source), Node 20→22, package-lock.json path trigger
- VitePress sidebar restructured into "入门" + "算法" groups, nav added "相关链接" dropdown
- README.md: fixed dead doc links (algorithms, benchmarks → guide/ paths)
- Upgraded Arithmetic benchmark from C++-only to cross-language (C++, Go, Rust)
- Overhauled Makefile with complete build/test/clean targets for all algorithms and languages
- Updated CI workflow: full build, test, and cross-language verification for all 4 algorithms × 3 languages
- Updated `go.work` to include `arithmetic/go` module
- Updated `range/rust/Cargo.toml` to include `rangecoder` CLI binary
- Updated `.gitignore` for new binaries and file extensions
- Updated `run_all_bench.py` to use unified cross-language benchmark scripts
- Updated README: algorithm table, project structure, Go version badge (1.21+), roadmap

### Fixed | 修复
- README badge URL: `docs.yml` → `pages.yml` (both EN & ZH READMEs)
- README algorithm table: "LZ77" → "Range Coder" (matched actual project directory)
- README project structure: `lz77/` → `range/`
- `.gitignore`: removed bare `Makefile` pattern that could shadow root Makefile
- Fixed Makefile `build-range` target (was running `go test` instead of building CLI)
- Fixed Makefile `clean` target (replaced fragile Python one-liner with proper `rm` commands)

## [1.0.0] - 2026-01-07

### Added | 新增

#### Algorithms | 算法
- Huffman encoding implementation in C++, Go, and Rust
- Arithmetic coding implementation in C++
- Range coder implementation in C++, Go, and Rust
- Run-Length Encoding (RLE) implementation in C++, Go, and Rust

#### Features | 功能
- Unified CLI interface across all implementations (`encode`/`decode` commands)
- Cross-language file format compatibility
- Benchmark scripts for performance comparison
- Test data generation scripts

#### Documentation | 文档
- Comprehensive README with bilingual (Chinese/English) content
- Algorithm comparison table
- Quick start guide
- Project structure documentation

#### Open Source Infrastructure | 开源基础设施
- MIT License
- Contributing guidelines (CONTRIBUTING.md)
- Code of Conduct (CODE_OF_CONDUCT.md)
- Security policy (SECURITY.md)
- GitHub Issue templates (bug report, feature request)
- GitHub Pull Request template
- GitHub Actions CI/CD pipeline
  - Multi-platform C++ builds (Ubuntu, macOS)
  - Go build and lint checks
  - Rust build and clippy checks
  - Encode/decode correctness verification
  - Required files check

### Technical Details | 技术细节

#### File Formats | 文件格式
- Huffman: Magic header `HFMN`, frequency table, bit-encoded data
- RLE: `(count, value)` pairs with 4-byte little-endian count

#### Build Requirements | 构建要求
- C++17 compatible compiler
- Go 1.21+
- Rust 1.70+
- Python 3.8+ (for scripts)

---

## Version History Summary | 版本历史摘要

| Version | Date | Highlights |
|---------|------|------------|
| 1.0.0 | 2026-01-07 | Initial release with 4 algorithms, 3 languages, full open source setup |

---

[Unreleased]: https://github.com/LessUp/encoding/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/LessUp/encoding/releases/tag/v1.0.0
