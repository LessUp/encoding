package main

import (
	"github.com/LessUp/compress-kit/algorithms/shared/go/cli"
	"rle"
)

type RLEProcessor struct{}

func (p *RLEProcessor) EncodeFile(inputPath, outputPath string) error {
	return rle.EncodeFile(inputPath, outputPath)
}

func (p *RLEProcessor) DecodeFile(inputPath, outputPath string) error {
	return rle.DecodeFile(inputPath, outputPath)
}

func main() {
	cli.Run("rle", &RLEProcessor{})
}
