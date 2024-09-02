package main

import (
	"fmt"
	"os/exec"
)

func ExtractAudio(videoPath, audioOutputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-q:a", "0", "-map", "a", audioOutputPath)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract audio: %w", err)
	}

	return nil

}

func main() {
	err := ExtractAudio("input_video.mp4", "output_audio.wav")
	if err != nil {
		fmt.Printf("Unable to extract audio %v\n", err)
	}
}
