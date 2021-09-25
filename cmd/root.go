package cmd

import (
	"fmt"
	"os"

	"github.com/honza/extract-subtitle/pkg/subtitle"
	"github.com/spf13/cobra"
)

var FfmpegBin string

var rootCmd = &cobra.Command{
	Use:   "extract-subtitle [video-file] [language] [output-file]",
	Short: "Extract embedded subtitles froma video file",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		err := subtitle.ExtractSubtitleStreamToFile(args[0], args[1], args[2], FfmpegBin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&FfmpegBin, "ffmpeg-bin", "ffmpeg", "Alternative path to ffmpeg")
}
