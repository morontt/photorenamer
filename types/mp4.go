package types

import (
	"errors"
	"os/exec"
	"regexp"
	"time"
)

var (
	regexpMediainfoOutput *regexp.Regexp
	regexpTimeformatA     *regexp.Regexp
	regexpTimeformatB     *regexp.Regexp
)

type Mpeg4File struct {
	baseMediaPart
}

func (mpeg *Mpeg4File) Extension() string {
	return "mp4"
}

func init() {
	regexpMediainfoOutput = regexp.MustCompile(`^RT=(.*)\s+ET=(.*)\n$`)
	// 2006-01-02T15:04:05-0700
	regexpTimeformatA = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}[\+-]\d{4}$`)
	// 2023-05-01 06:08:47 UTC
	regexpTimeformatB = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} \w{3,4}$`)
}

func (mpeg *Mpeg4File) ParseTime() error {
	var timeString string

	out, err := exec.Command("mediainfo", "--Output=General;RT=%Recorded_Date% ET=%Encoded_Date%", mpeg.filename).Output()
	if err != nil {
		return err
	}

	s := regexpMediainfoOutput.FindStringSubmatch(string(out))
	if len(s) < 3 {
		return errors.New("mp4: wrong mediainfo output format")
	}

	if len(s[1]) > 0 {
		timeString = s[1]
	} else if len(s[2]) > 0 {
		timeString = s[2]
	} else {
		return errors.New("mp4: DateTime is not present")
	}

	t, err := parseTimeString(timeString)
	if err != nil {
		return err
	}

	mpeg.dateTime = t.In(time.Local)

	return nil
}

func parseTimeString(timeString string) (time.Time, error) {
	var timeLayout string

	if regexpTimeformatA.MatchString(timeString) {
		timeLayout = "2006-01-02T15:04:05-0700"
	} else if regexpTimeformatB.MatchString(timeString) {
		timeLayout = "2006-01-02 15:04:05 MST"
	} else {
		return time.Time{}, errors.New("mp4: unknown DateTime format")
	}

	return time.Parse(timeLayout, timeString)
}
