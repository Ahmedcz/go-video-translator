package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Ahmedcz/go-video-translator/internal/recognition"
	"github.com/Ahmedcz/go-video-translator/internal/translation"
)

func main() {
	// Recognize the audio and write the file to a transcript
	transcriptFileName := "transcript.txt"
	file, err := os.Create(transcriptFileName)
	if err != nil {
		fmt.Printf("failed to create transcript file: %v\n", err)
	}
	defer file.Close()
	// call recognize with the file as the writer
	err = recognition.Recognize(file, "input_audio.wav")
	if err != nil {
		log.Fatalf("failed to recognize audio; %v\n", err)
	}
	// Open the transcript from file

	transcriptFile, err := os.Open(transcriptFileName)
	if err != nil {
		log.Fatalf("failed to open transcript file: %v", err)
	}

	defer transcriptFile.Close()

	// Read the content of the transcript file

	transcriptContent, err := io.ReadAll(transcriptFile)

	if err != nil {
		log.Fatalf("failed to read content of the transcript file: %v", err)
	}

	// Translate the transcript content

	err = translation.TranslateText(os.Stdout, " 108775865375", "fr", "en-US", string(transcriptContent))

	if err != nil {
		log.Fatalf("failed to translate transcript content: %v", err)
	}

}
