package command

import (
	"log"
	"os"
	"regexp"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func Run() {
	re := regexp.MustCompile(`(?i).jpe?g$`)

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fname := file.Name()
			if re.MatchString(fname) {
				time, err := decodeExif(fname)
				if err != nil {
					Error(fname, err)
				} else {
					Success(fname, time)
				}
			}
		}
	}
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
