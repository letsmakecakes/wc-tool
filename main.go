package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func main() {
	// Create a log file with error handling
	file, err := os.Create("./log/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	log.SetOutput(file)

	// Log a test entry
	log.Println("This is a test entry")

	// Handle command-line arguments
	args := os.Args
	log.Println("Arguments:", args)

	if len(args) < 4 {
		log.Error("Invalid usage: wc [-c|-w|-l] <filename.txt>")
		os.Exit(1)
	}

	if args[1] != "wc" {
		log.Error("Invalid command")
		os.Exit(1)
	}

	if args[2] != "-c" && args[2] != "-w" && args[2] != "-l" {
		log.Error("Invalid option")
		os.Exit(1)
	}

	filename := args[3]
	ext := strings.Split(filename, ".")
	if len(ext) < 2 || ext[1] != "txt" {
		log.Error("Unsupported file type or missing extension")
		os.Exit(1)
	}
}
