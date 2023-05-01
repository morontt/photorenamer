package types

import (
	"errors"
	"os/exec"
	"regexp"
)

var regexpMediainfoOutput *regexp.Regexp

type Mpeg4File struct {
	baseMediaPart
}

func (mpeg *Mpeg4File) Extension() string {
	return "mp4"
}

func init() {
	regexpMediainfoOutput = regexp.MustCompile(`^RT=(.*)\s+ET=(.*)\n$`)
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

	err = mpeg.setTimeByString(timeString)
	if err != nil {
		return err
	}

	return nil
}
