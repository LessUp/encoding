package main

import (
	"fmt"
	"os"

	"github.com/LessUp/compress-kit/algorithms/shared/go/cli"
	"github.com/LessUp/compress-kit/algorithms/shared/go/codec"
	"rangecoder"
)

type RangeProcessor struct{}

func (p *RangeProcessor) EncodeFile(inputPath, outputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("cannot open input file: %s: %w", inputPath, err)
	}

	encoded, err := codec.EncodeBuffer(rangecoder.NewStreamingEncoder(), data)
	if err != nil {
		return err
	}

	if err := os.WriteFile(outputPath, encoded, 0o644); err != nil {
		return fmt.Errorf("cannot open output file: %s: %w", outputPath, err)
	}

	return nil
}

func (p *RangeProcessor) DecodeFile(inputPath, outputPath string) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("cannot open input file: %s: %w", inputPath, err)
	}

	decoded, err := codec.DecodeBuffer(rangecoder.NewStreamingDecoder(), data)
	if err != nil {
		return err
	}

	if err := os.WriteFile(outputPath, decoded, 0o644); err != nil {
		return fmt.Errorf("cannot open output file: %s: %w", outputPath, err)
	}

	return nil
}

func main() {
	cli.Run("rangecoder", &RangeProcessor{})
}
