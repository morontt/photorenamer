package types

import (
	"errors"
)

type QuickTimeFile struct {
	baseMediaPart
}

func (mov *QuickTimeFile) Extension() string {
	return "mov"
}

func (mov *QuickTimeFile) ParseTime() error {
	track, err := extractGeneralTrack(mov.filename)
	if err != nil {
		return err
	}

	var timeString string
	var ok bool
	timeString, ok = track.Extra["com_apple_quicktime_creationdate"]
	if !ok {
		timeString, err = track.getTimeString()
		if err != nil {
			return err
		}
	}

	if len(timeString) == 0 {
		return errors.New("mov.go: DateTime is not present")
	}

	err = mov.setTimeByString(timeString)
	if err != nil {
		return err
	}

	return nil
}
