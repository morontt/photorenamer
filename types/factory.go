package types

import (
	"errors"
	"regexp"
)

var (
	regexpJpg       = regexp.MustCompile(`(?i)\.jpe?g$`)
	regexpMpeg4     = regexp.MustCompile(`(?i)\.mp4$`)
	regexpQuickTime = regexp.MustCompile(`(?i)\.mov$`)
)

func Support(filename string) bool {
	return regexpJpg.MatchString(filename) ||
		regexpMpeg4.MatchString(filename) ||
		regexpQuickTime.MatchString(filename)
}

func GetMediaFile(filename string) (MediaFile, error) {
	if regexpJpg.MatchString(filename) {
		return &JpegFile{
			baseMediaPart: baseMediaPart{filename: filename},
		}, nil
	}

	if regexpMpeg4.MatchString(filename) {
		return &Mpeg4File{
			baseMediaPart: baseMediaPart{filename: filename},
		}, nil
	}

	if regexpQuickTime.MatchString(filename) {
		return &QuickTimeFile{
			baseMediaPart: baseMediaPart{filename: filename},
		}, nil
	}

	return nil, errors.New("factory: unsupported file")
}
