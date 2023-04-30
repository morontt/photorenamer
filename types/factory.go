package types

import (
	"errors"
	"regexp"
)

var (
	regexpJpg *regexp.Regexp
)

func init() {
	regexpJpg = regexp.MustCompile(`(?i).jpe?g$`)
}

func Support(filename string) bool {
	return regexpJpg.MatchString(filename)
}

func GetMediaFile(filename string) (MediaFile, error) {
	if regexpJpg.MatchString(filename) {
		return &JpegFile{
			baseMediaPart: baseMediaPart{filename: filename},
		}, nil
	}

	return nil, errors.New("factory: unsupported file")
}
