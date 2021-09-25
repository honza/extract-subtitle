// extract-subtitle
// Copyright (C) 2021  Honza Pokorny <honza@pokorny.ca>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
