package types

import (
	"errors"
	"regexp"
	"time"
)

type MediaFile interface {
	Extension() string
	OriginalFilename() string
	ParseTime() error
	DateTime() string
}

type baseMediaPart struct {
	filename string
	dateTime time.Time
}

func (m *baseMediaPart) OriginalFilename() string {
	return m.filename
}

func (m *baseMediaPart) DateTime() string {
	return m.dateTime.Format("2006-01-02 15.04.05")
}

func (m *baseMediaPart) setTimeByString(timeString string) error {
	t, err := parseTimeString(timeString)
	if err != nil {
		return err
	}

	m.dateTime = t.In(time.Local)

	return nil
}

var (
	regexpTimeformatA *regexp.Regexp
	regexpTimeformatB *regexp.Regexp
)

func init() {
	// 2006-01-02T15:04:05-0700
	regexpTimeformatA = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[\+-]\d{4}$`)
	// 2023-05-01 06:08:47 UTC
	regexpTimeformatB = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} \w{3,4}$`)
}

func parseTimeString(timeString string) (time.Time, error) {
	var timeLayout string

	if regexpTimeformatA.MatchString(timeString) {
		timeLayout = "2006-01-02T15:04:05-0700"
	} else if regexpTimeformatB.MatchString(timeString) {
		timeLayout = "2006-01-02 15:04:05 MST"
	} else {
		return time.Time{}, errors.New("types.go: unknown DateTime format")
	}

	return time.Parse(timeLayout, timeString)
}
