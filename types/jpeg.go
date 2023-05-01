package types

import (
	"os"
	"time"

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

func decodeExif(fname string) (time.Time, error) {
	f, err := os.Open(fname)
	if err != nil {
		return time.Time{}, err
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	f.Close()
	if err != nil {
		return time.Time{}, err
	}

	return x.DateTime()
}
