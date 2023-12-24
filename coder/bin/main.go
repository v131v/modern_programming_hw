package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/v131v/modern_programming_hw/coder/src/coder"
)

var helpText = strings.Join([]string{"Usage:", "coder encode -i <input.file> -o <output.file>", "coder encode <input.file>", "coder decode -i <input.file> -o <output.file>", "coder decode <input.file>"}, "\n")

func main() {
	c := coder.Coder{}

	if len(os.Args) < 2 {
		fmt.Println(helpText)
		os.Exit(1)
	}

	if os.Args[1] == "encode" {
		encodeCmd := flag.NewFlagSet("encode", flag.ExitOnError)

		encodeCmd.StringVar(&c.InputFilePath, "i", "", "Input file")
		encodeCmd.StringVar(&c.OutputFilePath, "o", "", "Output file")

		err := encodeCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(helpText)
			os.Exit(1)
		}

		if c.InputFilePath == "" {

			if len(os.Args) < 3 {
				fmt.Println(helpText)
				os.Exit(1)
			}

			c.InputFilePath = os.Args[2]
			c.OutputFilePath = os.Args[2] + ".out"
		}

		c.EncodeFileToBase64()
	}

	if os.Args[1] == "decode" {
		decodeCmd := flag.NewFlagSet("decode", flag.ExitOnError)

		decodeCmd.StringVar(&c.InputFilePath, "i", "", "Input file")
		decodeCmd.StringVar(&c.OutputFilePath, "o", "", "Output file")

		err := decodeCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(helpText)
			os.Exit(1)
		}

		if c.InputFilePath == "" {

			if len(os.Args) < 3 {
				fmt.Println(helpText)
				os.Exit(1)
			}

			c.InputFilePath = os.Args[2]
			c.OutputFilePath = os.Args[2] + ".out"
		}

		c.DecodeFileToBase64()
	}
}
