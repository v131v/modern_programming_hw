package coder

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

type Coder struct {
	InputFilePath  string
	OutputFilePath string
}

func (c Coder) EncodeFileToBase64() error {

	inputFile, err := os.Open(c.InputFilePath)
	if err != nil {
		return fmt.Errorf("error opening input file: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(c.OutputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	base64Encoder := base64.NewEncoder(base64.StdEncoding, outputFile)

	_, err = io.Copy(base64Encoder, inputFile)
	if err != nil {
		return fmt.Errorf("error encoding file: %v", err)
	}

	base64Encoder.Close()

	fmt.Printf("File successfully encoded and saved to %s\n", c.OutputFilePath)
	return nil
}

func (c Coder) DecodeFileToBase64() error {

	inputFile, err := os.Open(c.InputFilePath)
	if err != nil {
		return fmt.Errorf("error opening input file: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(c.OutputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	base64Decoder := base64.NewDecoder(base64.StdEncoding, inputFile)

	_, err = io.Copy(outputFile, base64Decoder)
	if err != nil {
		return fmt.Errorf("error encoding file: %v", err)
	}

	fmt.Printf("File successfully encoded and saved to %s\n", c.OutputFilePath)
	return nil
}
