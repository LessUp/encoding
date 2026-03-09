package main

import (
	"fmt"
	"os"

	"rangecoder"
)

// Range coder CLI 封装。
// 读取整个文件到内存，调用 rangecoder 库执行编解码，写出结果。
// 文件格式与 C++/Rust 实现完全一致，支持交叉编解码验证。

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s encode|decode input output\n", os.Args[0])
		os.Exit(1)
	}

	mode := os.Args[1]
	inputPath := os.Args[2]
	outputPath := os.Args[3]

	switch mode {
	case "encode":
		data, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "无法读取输入文件: %v\n", err)
			os.Exit(1)
		}
		encoded, err := rangecoder.Encode(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "编码失败: %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile(outputPath, encoded, 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "无法写入输出文件: %v\n", err)
			os.Exit(1)
		}
	case "decode":
		data, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "无法读取输入文件: %v\n", err)
			os.Exit(1)
		}
		decoded, err := rangecoder.Decode(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "解码失败: %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile(outputPath, decoded, 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "无法写入输出文件: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "未知模式，应为 encode 或 decode")
		os.Exit(1)
	}
}
