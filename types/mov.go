package types

import (
	"encoding/json"
	"errors"
	"os/exec"
)

type QuickTimeFile struct {
	baseMediaPart
}

func (mov *QuickTimeFile) Extension() string {
	return "mov"
}

type jsonData struct {
	Media jsonTrack
}

type jsonTrack struct {
	Tracks []jsonTrackItem `json:"track"`
}

type jsonTrackItem struct {
	TypeItem string `json:"@type"`
	Extra    map[string]string
}

func (mov *QuickTimeFile) ParseTime() error {
	out, err := exec.Command("mediainfo", "--Output=JSON", mov.filename).Output()
	if err != nil {
		return err
	}

	var parts jsonData
	err = json.Unmarshal(out, &parts)
	if err != nil {
		return err
	}

	var timeString string
	var ok bool
	for _, item := range parts.Media.Tracks {
		if item.TypeItem == "General" {
			timeString, ok = item.Extra["com_apple_quicktime_creationdate"]
			if !ok {
				return errors.New("mov.go: DateTime is not present")
			}
		}
	}

	err = mov.setTimeByString(timeString)
	if err != nil {
		return err
	}

	return nil
}
