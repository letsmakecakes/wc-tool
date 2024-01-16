package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"wc"
)

func main() {
	// Define command-line flags
	option := flag.String("c", "-c", "Specify the counting option: -c, -l, -w, -m")
	flag.Parse()

	// Get the filename from the remaining command line arguments
	args := flag.Args()
	if len(args) == 0 {
		log.Println("Usage: go run main.go -<option> <filename>")
		os.Exit(1)
	}

	filename := args[1]

	// Perform word count based on the provided option
	result, err := wc.WC(*option, filename)
	if err != nil {
		log.Fatal(err)
	}

	// Print the result
	log.Println(result)
}
