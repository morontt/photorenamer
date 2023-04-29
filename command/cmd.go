package command

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

var files map[string]bool

func init() {
	files = make(map[string]bool)
}

func Run() {
	re := regexp.MustCompile(`(?i).jpe?g$`)

	dirEntries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dirEntries {
		if !file.IsDir() {
			fname := file.Name()
			if re.MatchString(fname) {
				files[fname] = true
			}
		}
	}

	for fname, value := range files {
		time, err := decodeExif(fname)
		if err != nil {
			Error(fname, err)
		} else {
			if value {
				move(fname, time)
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

func move(oldFilename, time string) {
	var i int = 1

	newFilename := time + ".jpg"
	if newFilename != oldFilename {
		_, ok := files[newFilename]
		for ok {
			if hash(oldFilename) == hash(newFilename) {
				Duplicate(oldFilename, newFilename)
				err := os.Remove(oldFilename)
				if err != nil {
					log.Fatal(err)
				}

				return
			}

			newFilename = fmt.Sprintf("%s(%d).jpg", time, i)
			i++
			_, ok = files[newFilename]
			if !ok {
				files[newFilename] = false
			}
		}

		err := os.Rename(oldFilename, newFilename)
		if err != nil {
			Error(oldFilename, err)
		} else {
			Success(oldFilename, newFilename)
		}
	}
}

func hash(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
