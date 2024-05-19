package types

type Mpeg4File struct {
	baseMediaPart
	mediaInfoPart
}

func (mpeg *Mpeg4File) Extension() string {
	return "mp4"
}

func (mpeg *Mpeg4File) ParseTime() error {
	timeString, err := mpeg.extractTimeString(mpeg.filename, mpeg.Extension())
	if err != nil {
		return err
	}

	err = mpeg.setTimeByString(timeString)
	if err != nil {
		return err
	}

	return nil
}
