package types

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	goexif "github.com/dsoprea/go-exif/v3"
)

type JpegFile struct {
	baseMediaPart
}

func (jpg *JpegFile) Extension() string {
	return "jpg"
}

func (jpg *JpegFile) ParseTime() error {
	dt, err := decodeExif(jpg.filename)
	if err != nil {
		return err
	}

	jpg.dateTime = dt

	return nil
}

func decodeExif(fname string) (time.Time, error) {
	f, err := os.Open(fname)
	if err != nil {
		return time.Time{}, err
	}

	fileData, err := io.ReadAll(f)
	if err != nil {
		return time.Time{}, err
	}

	rawExif, err := goexif.SearchAndExtractExifN(fileData, 0)
	if err != nil {
		return time.Time{}, err
	}

	entries, _, err := goexif.GetFlatExifData(rawExif, nil)
	if err != nil {
		return time.Time{}, err
	}

	return searchDateTime(entries)
}

func searchDateTime(exifTags []goexif.ExifTag) (time.Time, error) {
	var dt time.Time
	tag, err := findTag("DateTimeOriginal", exifTags)
	if err != nil {
		tag, err = findTag("DateTime", exifTags)
		if err != nil {
			return dt, err
		}
	}

	return time.Parse("2006:01:02 15:04:05", tag.Formatted)
}

func findTag(tagName string, exifTags []goexif.ExifTag) (goexif.ExifTag, error) {
	for _, tag := range exifTags {
		if tag.TagName == tagName {
			return tag, nil
		}
	}

	return goexif.ExifTag{}, errors.New(fmt.Sprintf("exif: tag %s is not present", tagName))
}
