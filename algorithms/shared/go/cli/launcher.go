package cli

import (
	"fmt"
	"os"
)

type FileProcessor interface {
	EncodeFile(inputPath, outputPath string) error
	DecodeFile(inputPath, outputPath string) error
}

type BufferProcessor interface {
	NewEncoder() interface{ Process([]byte) ([]byte, error) }
	NewDecoder() interface{ Process([]byte) ([]byte, error) }
}

func Run(name string, processor FileProcessor) {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s encode|decode input output\n", os.Args[0])
		os.Exit(1)
	}

	mode := os.Args[1]
	inputPath := os.Args[2]
	outputPath := os.Args[3]

	var err error
	switch mode {
	case "encode":
		err = processor.EncodeFile(inputPath, outputPath)
	case "decode":
		err = processor.DecodeFile(inputPath, outputPath)
	default:
		fmt.Fprintln(os.Stderr, "unknown mode, expected encode or decode")
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
