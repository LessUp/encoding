package main

import (
	"arithmetic"
	"github.com/LessUp/compress-kit/algorithms/shared/go/cli"
)

type ArithmeticProcessor struct{}

func (p *ArithmeticProcessor) EncodeFile(inputPath, outputPath string) error {
	return arithmetic.EncodeFile(inputPath, outputPath)
}

func (p *ArithmeticProcessor) DecodeFile(inputPath, outputPath string) error {
	return arithmetic.DecodeFile(inputPath, outputPath)
}

func main() {
	cli.Run("arithmetic", &ArithmeticProcessor{})
}
