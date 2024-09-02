package streaming

import (
	"fmt"
	"io"
	"os"
)

// Processing

func ProcessInput(inputPath string) (*os.File, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("Error opening %s: %v\n", inputPath, err)
	}

	return file, nil
}

// Decoding

func DecodeData(file *os.File) ([]byte, error) {
	var data []byte
	_, err := file.Read(data)
	if err != nil && err != io.EOF {
		fmt.Errorf("Failed to decode data: %v\n", err)
	}
	return nil, data
}

// Overlaying

func ApplyOverlay(data []byte, overlay []byte) []byte {
	return append(data, overlay...)
}

// Streaming

func StreamOutput(outputPath string, data []byte) error {
	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Errorf("Failed to create file %s: %v\n", err)
	}

	_, err := file.Write(data)
	if err != nil {
		fmt.Errorf("Failed to write data to file: %v\n", err)
	}

	return nil
}

func ProcessAndStream(inputPath string, overlayPath string, outputPath string) error {
	//ProcessInput

	inputFile, err := ProcessInput(inputPath)

	if err != nil {
		return err
	}

	defer inputFile.Close()

	// DecodeData

	err, data := DecodeData(inputFile)

	if err != nil {
		return err
	}
	// ApplyOverlay

	err, overlayFile := ProcessInput(overlayPath)
	if err != nil {
		return err
	}

	defer overlayFile.Close()

	err, overlayData := DecodeData(overlayFile)
	if err != nil {
		return err
	}

	resultData := ApplyOverlay(data, overlayData)

	//Streaming
	StreamOutput(outputPath, resultData)

}
