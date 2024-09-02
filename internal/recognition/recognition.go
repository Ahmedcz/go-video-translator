package recognition

import (
	"context"
	"fmt"
	"io"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
	"google.golang.org/api/option"
)

func Recognize(w io.Writer, file string) error {
	ctx := context.Background()
	client, err := speech.NewClient(ctx, option.WithCredentialsFile("C:\\Users\\Lenovo\\Desktop\\Video Translator_Go\\internal\\1_translation\\real-time-video-translator-e5a1c98884ae.json"))
	if err != nil {
		return err
	}
	defer client.Close()

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	// Send the contents of the audio file with the encoding and sample rate information to be recognized.

	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: data},
		},
	})

	// Print the results.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Fprintf(w, "\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}
	return nil
}
