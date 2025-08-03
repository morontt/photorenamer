package types

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type MediaFile interface {
	Extension() string
	OriginalFilename() string
	TimeBasedFilename() string
	ParseTime() error
	DateTime() string
	Hash() string
}

type MediaFiles []MediaFile

var (
	// 2006-01-02T15:04:05-0700
	regexpTimeformatA = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[\+-]\d{4}$`)
	// 2023-05-01 06:08:47 UTC
	regexpTimeformatB = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} \w{3,4}$`)
	// 2025-07-17 20:33:27+03:00
	regexpTimeformatC = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}[\+-]\d{2}:\d{2}$`)
)

func parseTimeString(timeString string) (time.Time, error) {
	var timeLayout string

	if regexpTimeformatA.MatchString(timeString) {
		timeLayout = "2006-01-02T15:04:05-0700"
	} else if regexpTimeformatB.MatchString(timeString) {
		timeLayout = "2006-01-02 15:04:05 MST"
	} else if regexpTimeformatC.MatchString(timeString) {
		timeLayout = "2006-01-02 15:04:05-07:00"
	} else {
		return time.Time{}, errors.New("types.go: unknown DateTime format")
	}

	return time.Parse(timeLayout, timeString)
}

func (mf MediaFiles) Len() int {
	return len(mf)
}

func (mf MediaFiles) Swap(i, j int) {
	mf[i], mf[j] = mf[j], mf[i]
}

func (mf MediaFiles) Less(i, j int) bool {
	a, b := mf[i], mf[j]

	cmp := strings.Compare(
		strings.ToLower(a.TimeBasedFilename()),
		strings.ToLower(b.TimeBasedFilename()),
	)
	if cmp == 0 {
		cmp = strings.Compare(a.Hash(), b.Hash())
	}

	return cmp < 0
}
