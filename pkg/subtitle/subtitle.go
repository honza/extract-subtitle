package subtitle

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func listStreams(ffmpegBin string, filename string) ([]string, error) {
	streams := []string{}
	cmd := exec.Command(ffmpegBin, "-i", filename)
	b, _ := cmd.CombinedOutput()
	s := string(b)

	lines := strings.Split(s, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "  Stream") {
			streams = append(streams, line)
		}
	}

	if len(streams) == 0 {
		return streams, errors.New("No streams found")
	}

	return streams, nil

}

func getSubtitleStreamNumber(ffmpegBin string, filename string, language string) (string, error) {
	streams, err := listStreams(ffmpegBin, filename)
	if err != nil {
		return "", err
	}

	p := regexp.MustCompile(`#\d:(\d+)`)

	for _, stream := range streams {
		if strings.Contains(stream, "Subtitle") {
			if strings.Contains(stream, language) {
				fmt.Println(stream)
				m := p.FindStringSubmatch(stream)
				if len(m) != 0 {
					return m[1], nil
				}
			}
		}
	}

	return "", errors.New("Subtitle with specified language not found")
}

func ExtractSubtitleStreamToFile(source string, language string, output, ffmpegBin string) error {
	number, err := getSubtitleStreamNumber(ffmpegBin, source, language)
	if err != nil {
		return err
	}

	mapping := fmt.Sprintf("0:%s", number)
	cmd := exec.Command(ffmpegBin, "-i", source, "-map", mapping, output)
	return cmd.Run()
}
