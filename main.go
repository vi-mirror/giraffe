package main

import (
	"flag"
	"log"

	"github.com/tatthien/giraffe/cmd"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		cmd.Build()
		return
	}

	command := flag.Arg(0)
	switch command {
	case "serve":
		cmd.Serve()
	default:
		log.Println("Unknown command:", command)
	}
}
