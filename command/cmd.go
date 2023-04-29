package command

import (
	"fmt"
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
				fmt.Println(fname)
				decodeExif(fname)
			}
		}
	}
}

func decodeExif(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	tm, err := x.DateTime()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Taken: ", tm)
}
