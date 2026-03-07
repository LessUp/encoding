.PHONY: build build-huffman build-arithmetic build-range build-rle test test-huffman-go test-range-go test-rle-go test-huffman-rust test-range-rust test-rle-rust test-data bench clean

build: build-huffman build-arithmetic build-range build-rle

build-huffman:
	g++ -std=c++17 -O2 huffman/cpp/main.cpp -o huffman/cpp/huffman_cpp
	go build -o huffman/go/huffman_go ./huffman/go
	rustc -O huffman/rust/main.rs -o huffman/rust/huffman_rust

build-arithmetic:
	g++ -std=c++17 -O2 arithmetic/cpp/main.cpp -o arithmetic/cpp/arithmetic_cpp

build-range:
	g++ -std=c++17 -O2 range/cpp/main.cpp -o range/cpp/rangecoder_cpp
	go test ./range/go/...
	cargo build --manifest-path range/rust/Cargo.toml --release

build-rle:
	g++ -std=c++17 -O2 rle/cpp/main.cpp -o rle/cpp/rle_cpp
	go build -o rle/go/rle_go ./rle/go
	rustc -O rle/rust/main.rs -o rle/rust/rle_rust

test: test-data test-huffman-go test-range-go test-rle-go test-huffman-rust test-range-rust test-rle-rust

test-huffman-go:
	go test ./huffman/go/...

test-range-go:
	go test ./range/go/...

test-rle-go:
	go test ./rle/go/...

test-huffman-rust:
	rustc --test huffman/rust/main.rs -o huffman/rust/huffman_rust_test
	./huffman/rust/huffman_rust_test

test-range-rust:
	cargo test --manifest-path range/rust/Cargo.toml

test-rle-rust:
	rustc --test rle/rust/main.rs -o rle/rust/rle_rust_test
	./rle/rust/rle_rust_test

test-data:
	python tests/gen_testdata.py

bench: test-data
	python scripts/run_all_bench.py

clean:
	python -c "from pathlib import Path; import shutil; root = Path('.'); patterns = ['reports', 'tests/data']; files = ['huffman/cpp/huffman_cpp', 'huffman/go/huffman_go', 'huffman/rust/huffman_rust', 'huffman/rust/huffman_rust_test', 'arithmetic/cpp/arithmetic_cpp', 'range/cpp/rangecoder_cpp', 'rle/cpp/rle_cpp', 'rle/go/rle_go', 'rle/rust/rle_rust', 'rle/rust/rle_rust_test']; [shutil.rmtree(root / p, ignore_errors=True) for p in patterns]; [(root / f).unlink(missing_ok=True) for f in files]"
