package command

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"xelbot.com/renamer/types"
)

var (
	files      map[string]bool
	regexpSkip = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}\.\d{2}\.\d{2}\.`)
)

func init() {
	files = make(map[string]bool)
}

func Run() {
	dirEntries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dirEntries {
		if !file.IsDir() {
			fname := file.Name()
			if types.Support(fname) && !skip(fname) {
				files[fname] = true
			}
		}
	}

	for fname, value := range files {
		if value {
			media, err := types.GetMediaFile(fname)
			if err != nil {
				Error(fname, err)
			} else {
				err := media.ParseTime()
				if err != nil {
					Error(fname, err)
				} else {
					move(media)
				}
			}
		}
	}
}

func skip(filename string) (result bool) {
	if regexpSkip.MatchString(filename) {
		result = true
		Skipped(filename)
	}

	return
}

func move(media types.MediaFile) {
	var i int = 1

	oldFilename := media.OriginalFilename()
	time := media.DateTime()

	newFilename := time + "." + media.Extension()
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

			newFilename = fmt.Sprintf("%s(%d).%s", time, i, media.Extension())
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
