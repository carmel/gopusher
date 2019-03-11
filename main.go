package main

import (
	"flag"
	"gopusher/comet"
	"gopusher/log"

	"github.com/joho/godotenv"
)

func main() {
	filename := getArgs()

	log.Info("Load config file: %s", *filename)
	if err := godotenv.Load(*filename); err != nil {
		panic(err)
	}

	comet.Run()
}

func getArgs() (filename *string) {
	filename = flag.String("c", "./.env", "set config file path")
	flag.Parse()

	return
}
