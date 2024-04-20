package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"xelbot.com/renamer/command"
)

type parameters struct {
	Recursive bool `short:"r" long:"recursive" description:"Rename files in directories recursively"`
}

func main() {
	args := new(parameters)
	_, err := flags.Parse(args)
	if err != nil {
		os.Exit(-1)
	}

	command.Run()
}
