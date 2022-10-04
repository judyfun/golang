package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
)

func main() {

	var cfg Config
	app := kingpin.New("goBegin", "test kingpin function")
	app.Flag("goBegin.name", "get goBegin name").Envar("nine_mel_name").StringVar(&cfg.Name)

	osArgs := os.Args[1:]
	command, err := app.Parse(osArgs)

	if err != nil {
		log.Fatalf("Failed to parse the flags, command: %v", command)
	}

}
