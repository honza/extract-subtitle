extract-subtitle
================

Many video files come with embedded subtitles, and many video players can't
access those.  This tiny utility extracts the subtitle from the video file into
a separate subtitle file which is much more widely supported.

## Dependencies

Golang, and ffmpeg

## Install

```
$ go get github.com/honza/extract-subtitle
```

## Usage

```
$ extract-subtitle --help
Extract embedded subtitles froma video file

Usage:
  extract-subtitle [video-file] [language] [output-file] [flags]

Flags:
      --ffmpeg-bin string   Alternative path to ffmpeg (default "ffmpeg")
  -h, --help                help for extract-subtitle
```

## License

GPLv3
