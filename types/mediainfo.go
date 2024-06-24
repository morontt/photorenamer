package types

import (
	"encoding/json"
	"errors"
	"os/exec"
)

type jsonData struct {
	Media jsonTrack
}

type jsonTrack struct {
	Tracks []jsonTrackItem `json:"track"`
}

type jsonTrackItem struct {
	TypeItem     string `json:"@type"`
	RecordedDate string `json:"Recorded_Date"`
	EncodedDate  string `json:"Encoded_Date"`
	Extra        map[string]string
}

func (j *jsonTrackItem) getTimeString() (string, error) {
	var timeString string
	if len(j.RecordedDate) > 0 {
		timeString = j.RecordedDate
	} else if len(j.EncodedDate) > 0 {
		timeString = j.EncodedDate
	} else {
		return "", errors.New("jsonTrackItem: DateTime is not present")
	}

	return timeString, nil
}

func extractGeneralTrack(filename string) (jsonTrackItem, error) {
	var track jsonTrackItem

	out, err := exec.Command("mediainfo", "--Output=JSON", filename).Output()
	if err != nil {
		return track, err
	}

	var parts jsonData
	err = json.Unmarshal(out, &parts)
	if err != nil {
		return track, err
	}

	for _, item := range parts.Media.Tracks {
		if item.TypeItem == "General" {
			return item, nil
		}
	}

	return track, errors.New("no general track found")
}
