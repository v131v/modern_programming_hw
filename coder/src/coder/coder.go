package coder

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func calculateMD5(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}

type Coder struct {
	InputFilePath  string
	OutputFilePath string
}

func (c Coder) EncodeFileToBase64() ([]byte, error) {

	inputFile, err := os.Open(c.InputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening input file: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(c.OutputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	base64Encoder := base64.NewEncoder(base64.StdEncoding, outputFile)

	_, err = io.Copy(base64Encoder, inputFile)
	if err != nil {
		return nil, fmt.Errorf("error encoding file: %v", err)
	}

	base64Encoder.Close()

	hash := md5.New()
	io.Copy(hash, inputFile)
	hashSum := hash.Sum([]byte{})

	fmt.Printf("File successfully encoded and saved to %s\n", c.OutputFilePath)
	return hashSum, nil
}

func (c Coder) DecodeFileToBase64() ([]byte, error) {

	inputFile, err := os.Open(c.InputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening input file: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(c.OutputFilePath)
	if err != nil {
		return nil, fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	base64Decoder := base64.NewDecoder(base64.StdEncoding, inputFile)

	_, err = io.Copy(outputFile, base64Decoder)
	if err != nil {
		return nil, fmt.Errorf("error encoding file: %v", err)
	}

	hash := md5.New()
	io.Copy(hash, outputFile)
	hashSum := hash.Sum([]byte{})

	fmt.Printf("File successfully encoded and saved to %s\n", c.OutputFilePath)
	return hashSum, nil
}
