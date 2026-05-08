package main

import (
	"github.com/LessUp/compress-kit/algorithms/shared/go/cli"
	"huffman"
)

type HuffmanProcessor struct{}

func (p *HuffmanProcessor) EncodeFile(inputPath, outputPath string) error {
	return huffman.EncodeFile(inputPath, outputPath)
}

func (p *HuffmanProcessor) DecodeFile(inputPath, outputPath string) error {
	return huffman.DecodeFile(inputPath, outputPath)
}

func main() {
	cli.Run("huffman", &HuffmanProcessor{})
}
