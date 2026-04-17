# Project Structure

This guide explains the project organization, file formats, and conventions used across all implementations.

## Directory Layout

```
encoding/
├── algorithms/huffman/              # Huffman coding implementation
│   ├── cpp/              #   C++ single-file implementation
│   ├── go/               #   Go module (go.mod)
│   ├── rust/             #   Rust implementation
│   └── benchmark/        #   Cross-language benchmark scripts
├── algorithms/arithmetic/           # Arithmetic coding implementation
│   ├── cpp/              #   C++ single-file implementation
│   ├── go/               #   Go implementation
│   ├── rust/             #   Rust implementation
│   └── benchmark/        #   Cross-language benchmark
├── algorithms/range/                # Range coder implementation
│   ├── cpp/              #   C++ single-file implementation
│   ├── go/               #   Go library + CLI
│   ├── rust/             #   Rust library crate + CLI
│   └── benchmark/        #   Cross-language benchmark
├── algorithms/rle/                  # Run-length encoding
│   ├── cpp/              #   C++ single-file implementation
│   ├── go/               #   Go implementation
│   ├── rust/             #   Rust implementation
│   └── benchmark/        #   Cross-language benchmark
├── tests/                # Test data generation
│   ├── gen_testdata.py   #   Generate benchmark test files
│   └── data/             #   Generated test data
├── docs/                 # Documentation site (VitePress)
│   ├── .vitepress/       #   VitePress configuration
│   ├── en/               #   English documentation
│   ├── zh/               #   Chinese documentation
│   └── public/           #   Static assets (logo, etc.)
├── .github/workflows/    # GitHub Actions CI/CD
├── Makefile              # Build, test, and benchmark entry point
├── package.json          # npm scripts for docs
└── go.work               # Go workspace (multi-module)
```

## Language Implementation Standards

| Language | Version | Build Method | Characteristics |
|----------|---------|--------------|-----------------|
| **C++** | C++17 | `g++ -std=c++17 -O2` | Single file, zero dependencies |
| **Go** | 1.21+ | Go modules (`go.mod` + `cmd/`) | All implementations provide library API + CLI |
| **Rust** | 1.70+ | Cargo / `rustc` | Range coder provides library crate |

## Unified CLI Interface

All implementations follow the same command-line pattern:

```bash
<algorithm>_<lang> encode <input_file> <output_file>
<algorithm>_<lang> decode <input_file> <output_file>
```

### Binary Names

| Algorithm | C++ | Go | Rust |
|-----------|-----|-----|------|
| Huffman | `huffman_cpp` | `huffman_go` | `huffman_rust` |
| Arithmetic | `arithmetic_cpp` | `arithmetic_go` | `arithmetic_rust` |
| Range Coder | `rangecoder_cpp` | `rangecoder_go` | `rangecoder` (cargo) |
| RLE | `rle_cpp` | `rle_go` | `rle_rust` |

### Usage Examples

```bash
# Encode with Huffman (C++)
./huffman_cpp encode document.txt document.huf

# Decode with Range Coder (Go)
./rangecoder_go decode data.rcnc data.bin

# Encode with RLE (Rust)
./rle_rust encode bitmap.raw bitmap.rle
```

## File Format Compatibility

All language implementations of the same algorithm use **identical binary formats**.

### Format Summary

| Algorithm | Magic Header | Extension | Structure |
|-----------|--------------|-----------|-----------|
| Huffman | `HFMN` | `.huf` | Magic + frequency table + bit stream |
| Arithmetic | `AENC` | `.aenc` | Magic + frequency table + bit stream |
| Range Coder | `RCNC` | `.rcnc` | Magic + frequency table + byte stream |
| RLE | None | `.rle` | (count: 4B LE, value: 1B) pairs |

### Cross-Language Verification

| Encode ↓ / Decode → | C++ | Go | Rust |
|---------------------|-----|-----|------|
| C++ | ✓ | ✓ | ✓ |
| Go | ✓ | ✓ | ✓ |
| Rust | ✓ | ✓ | ✓ |

Any combination works: **C++ ↔ Go ↔ Rust**

## CI/CD Pipeline

### Workflows

| Workflow | File | Trigger | Purpose |
|----------|------|---------|---------|
| **CI** | `.github/workflows/ci.yml` | Push / PR | Build, test, correctness verification |
| **Pages** | `.github/workflows/pages.yml` | `docs/` changes | Deploy documentation |

### CI Job Matrix

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  build-cpp  │     │  build-go   │     │ build-rust  │
│  ├ Ubuntu   │     │  Ubuntu     │     │  Ubuntu     │
│  └ macOS    │     │             │     │             │
└──────┬──────┘     └──────┬──────┘     └──────┬──────┘
       │                   │                   │
       └───────────────────┼───────────────────┘
                           │
                    ┌──────▼──────┐
                    │ correctness │
                    │   tests     │
                    │  (Python)   │
                    └─────────────┘
```

### CI Checks

1. **Build Jobs**: Compile all implementations on Ubuntu and macOS (C++)
2. **Test Jobs**: Run Go `go test` and Rust `cargo test`
3. **Lint Jobs**: Go vet, Rust clippy
4. **Correctness Jobs**: Cross-language encode/decode verification

## Security Considerations

### Input/Output Size Limits

All implementations enforce the following limits:

| Limit | Value | Purpose |
|-------|-------|---------|
| Max input size | 4 GiB | Prevent frequency overflow |
| Max output size | 1 GiB | Prevent decompression bombs |

These limits are applied before processing begins to prevent:
- Integer overflow attacks
- Decompression bomb attacks
- Excessive memory usage

## Build System Details

### Makefile Targets

| Target | Description |
|--------|-------------|
| `build` | Build all implementations |
| `build-huffman` | Build Huffman (C++, Go, Rust) |
| `build-arithmetic` | Build Arithmetic (C++, Go, Rust) |
| `build-range` | Build Range Coder (C++, Go, Rust) |
| `build-rle` | Build RLE (C++, Go, Rust) |
| `test` | Run all Go and Rust unit tests |
| `bench` | Generate test data and run all benchmarks |
| `test-data` | Generate test data only |
| `clean` | Remove all build artifacts and reports |

### Go Workspace

The project uses Go workspaces to manage multiple modules:

```go
// go.work
go 1.21

use (
    ./algorithms/huffman/go
    ./algorithms/arithmetic/go
    ./algorithms/range/go
    ./algorithms/rle/go
)
```

This allows:
- Cross-module dependencies
- Unified build commands
- IDE support for all Go modules

## Documentation Site

The documentation is built with [VitePress](https://vitepress.dev/):

```bash
# Install dependencies
npm install

# Start development server
npm run docs:dev

# Build for production
npm run docs:build
```

The built site is in `docs/.vitepress/dist/` and deployed to GitHub Pages.

---

## Related Documentation

- [Getting Started](/en/guide/getting-started) - Setup and basic usage
- [Algorithms Guide](/en/guide/algorithms) - Algorithm explanations
- [GitHub Repository](https://github.com/LessUp/encoding) - Source code
