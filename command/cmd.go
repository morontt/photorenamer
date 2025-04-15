package command

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"xelbot.com/renamer/types"
)

var (
	existingFiles map[string]bool
	regexpSkip    = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}\.\d{2}\.\d{2}(-\d+)?\.`)
)

func init() {
	existingFiles = make(map[string]bool)
}

func Run() {
	dirEntries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	mediaFiles := make(types.MediaFiles, 0)
	for _, file := range dirEntries {
		if !file.IsDir() {
			fname := file.Name()
			if types.Support(fname) {
				existingFiles[fname] = true
				media, err := types.GetMediaFile(fname)
				if err != nil {
					Error(fname, err)
				} else {
					if !skip(fname) {
						err = media.ParseTime()
						if err != nil {
							Error(fname, err)
						}
					}
					mediaFiles = append(mediaFiles, media)
				}
			}
		}
	}

	var newFilename string

	sort.Sort(mediaFiles)
	for idx, media := range mediaFiles {
		if idx > 0 {
			if mediaFiles[idx-1].TimeBasedFilename() == media.TimeBasedFilename() {
				if mediaFiles[idx-1].Hash() == media.Hash() {
					Duplicate(media.OriginalFilename(), mediaFiles[idx-1].OriginalFilename())
					err = os.Remove(media.OriginalFilename())
					if err != nil {
						log.Fatal(err)
					}
					delete(existingFiles, media.OriginalFilename())

					continue
				}
			}
		}

		if media.OriginalFilename() != media.TimeBasedFilename() {
			newFilename, err = move(media)
			if err != nil {
				Error(media.OriginalFilename(), err)
			} else {
				Success(media.OriginalFilename(), newFilename)
			}
		} else {
			Skipped(media.OriginalFilename())
		}
	}
}

func skip(filename string) (result bool) {
	return regexpSkip.MatchString(filename)
}

func move(media types.MediaFile) (string, error) {
	var i int
	var newFilename string

	last := strings.LastIndex(media.TimeBasedFilename(), ".")
	baseName := media.TimeBasedFilename()[:last]
	ext := media.Extension()

	i = -1
	for {
		i++
		if i == 0 {
			newFilename = baseName + "." + ext
		} else {
			newFilename = fmt.Sprintf("%s-%d.%s", baseName, i, ext)
		}

		if _, ok := existingFiles[newFilename]; ok {
			continue
		}

		_, err := os.Stat(newFilename)
		if errors.Is(err, fs.ErrNotExist) {
			err = os.Rename(media.OriginalFilename(), newFilename)
			if err != nil {
				return "", err
			}

			existingFiles[newFilename] = true

			break
		} else if err != nil {
			return "", err
		}
	}

	return newFilename, nil
}
