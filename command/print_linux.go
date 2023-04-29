package command

import "fmt"

const (
	Blue    = "\033[1;34m"
	Red     = "\033[1;31m"
	Magenta = "\033[1;35m"
	Reset   = "\033[0m"
)

func Error(filename string, err error) {
	fmt.Printf("%s%s%s %s\n", Red, filename, Reset, err.Error())
}

func Success(oldFilename, newFilename string) {
	fmt.Printf("%s%s%s renamed to %s\n", Blue, oldFilename, Reset, newFilename)
}

func Duplicate(oldFilename, newFilename string) {
	fmt.Printf("%s%s%s is a duplicate of file %s\n", Magenta, oldFilename, Reset, newFilename)
}
