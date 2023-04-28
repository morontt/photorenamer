package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

func main() {
	fname := "IMG_0004.JPEG"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	exif.RegisterParsers(mknote.All...)

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)
}
