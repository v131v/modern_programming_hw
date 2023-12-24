package coder

import (
	"bytes"
	"os"
	"testing"
)

func prepareEnv(inputData []byte) (string, string, error) {
	inputFilePath := "test_input.txt"
	outputFilePath := "test_output_base64.txt"

	err := os.WriteFile(inputFilePath, inputData, 0644)
	if err != nil {
		return "", "", err
	}

	return inputFilePath, outputFilePath, nil
}

func TestEncodeDecodeWithDigits(t *testing.T) {
	inputData := []byte("test text 123")

	inputFilePath, outputFilePath, err := prepareEnv(inputData)
	if err != nil {
		t.Fatalf("error creating test input file: %v", err)
	}
	defer os.Remove(inputFilePath)
	defer os.Remove(outputFilePath)

	c := Coder{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}

	err = c.EncodeFileToBase64()
	if err != nil {
		t.Fatalf("error encoding file to base64: %v", err)
	}

	c.InputFilePath, c.OutputFilePath = c.OutputFilePath, c.InputFilePath

	err = c.DecodeFileToBase64()
	if err != nil {
		t.Fatalf("error decoding file from base64: %v", err)
	}

	decodedData, err := os.ReadFile(inputFilePath)
	if err != nil {
		t.Fatalf("error reading decoded file: %v", err)
	}

	if !bytes.Equal(decodedData, inputData) {
		t.Errorf("decoded data doesn't match the original input data \nExpected: %v\nRecieved: %v", inputData, decodedData)
	}
}

func TestEncodeDecodeWithSymbols(t *testing.T) {
	inputData := []byte("some other text %(^(%__+^))")

	inputFilePath, outputFilePath, err := prepareEnv(inputData)
	if err != nil {
		t.Fatalf("error creating test input file: %v", err)
	}
	defer os.Remove(inputFilePath)
	defer os.Remove(outputFilePath)

	c := Coder{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}

	err = c.EncodeFileToBase64()
	if err != nil {
		t.Fatalf("error encoding file to base64: %v", err)
	}

	c.InputFilePath, c.OutputFilePath = c.OutputFilePath, c.InputFilePath

	err = c.DecodeFileToBase64()
	if err != nil {
		t.Fatalf("error decoding file from base64: %v", err)
	}

	decodedData, err := os.ReadFile(inputFilePath)
	if err != nil {
		t.Fatalf("error reading decoded file: %v", err)
	}

	if !bytes.Equal(decodedData, inputData) {
		t.Errorf("decoded data doesn't match the original input data \nExpected: %v\nRecieved: %v", inputData, decodedData)
	}
}
