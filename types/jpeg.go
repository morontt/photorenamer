package types

import (
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

type JpegFile struct {
	baseMediaPart
}

func (jpg *JpegFile) Extension() string {
	return "jpg"
}

func (jpg *JpegFile) ParseTime() error {
	time, err := decodeExif(jpg.filename)
	if err != nil {
		return err
	}

	jpg.dateTime = time

	return nil
}

func decodeExif(fname string) (string, error) {
	f, err := os.Open(fname)
	if err != nil {
		return "", err
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	f.Close()
	if err != nil {
		return "", err
	}

	tm, err := x.DateTime()
	if err != nil {
		return "", err
	}

	return tm.Format("2006-01-02 15.04.05"), nil
}
