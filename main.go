package main

import (
	"os"

	"vineelsai.com/checkout/project"
	"vineelsai.com/checkout/utils"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("not enough arguments")
	}

	switch args[0] {
	case "init":
		if len(args) == 1 {
			args = append(args, project.PromptForName())
		}
		project.Init(args[1])
	case "deinit":
		if len(args) == 2 {
			project.DeInit(args[1], utils.DefaultProjectSource)
		} else {
			project.DeInit(args[1], args[2])
		}
	}
}
