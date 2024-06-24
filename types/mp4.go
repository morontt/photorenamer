package types

type Mpeg4File struct {
	baseMediaPart
}

func (mpeg *Mpeg4File) Extension() string {
	return "mp4"
}

func (mpeg *Mpeg4File) ParseTime() error {
	track, err := extractGeneralTrack(mpeg.filename)
	if err != nil {
		return err
	}

	timeString, err := track.getTimeString()
	if err != nil {
		return err
	}

	err = mpeg.setTimeByString(timeString)
	if err != nil {
		return err
	}

	return nil
}
