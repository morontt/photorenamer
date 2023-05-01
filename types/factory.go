package types

import (
	"errors"
	"regexp"
)

var (
	regexpJpg   *regexp.Regexp
	regexpMpeg4 *regexp.Regexp
)

func init() {
	regexpJpg = regexp.MustCompile(`(?i)\.jpe?g$`)
	regexpMpeg4 = regexp.MustCompile(`(?i)\.mp4$`)
}

func Support(filename string) bool {
	return regexpJpg.MatchString(filename) ||
		regexpMpeg4.MatchString(filename)
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

	return nil, errors.New("factory: unsupported file")
}
