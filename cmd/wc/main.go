package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"wc-tool/internal/counter"
	"wc-tool/pkg/utils"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Add version flag
	versionFlag := flag.Bool("v", false, "print version information")

	// Define flags
	bytesFlag := flag.Bool("c", false, "print the byte counts")
	linesFlag := flag.Bool("l", false, "print the newline counts")
	wordsFlag := flag.Bool("w", false, "print the word counts")
	charsFlag := flag.Bool("m", false, "print the character counts")
	flag.Parse()

	// Handle version flag
	if *versionFlag {
		fmt.Printf("ccwc version %s, commit %s, built at %s\n", version, commit, date)
		return
	}

	// Set default flags if none are provided
	setDefaultFlags(bytesFlag, linesFlag, wordsFlag, charsFlag)

	// Determine input source (file or stdin)
	input, filename := getInputSource(flag.Args())
	defer func(input io.ReadCloser) {
		if input != os.Stdin {
			input.Close()
		}
	}(input)

	// Create counter and process input
	c := counter.New()
	err := c.Process(input)
	utils.CheckError(err, "Error processing input")

	// Collect results and format results
	results := collectResults(c, *linesFlag, *wordsFlag, *bytesFlag, *charsFlag)
	fmt.Print(utils.FormatOutput(results, filename))
}

// setDefaultFlags sets default flags if none are provided
func setDefaultFlags(bytesFlag, linesFlag, wordsFlag, charsFlag *bool) {
	if !*bytesFlag && !*linesFlag && !*wordsFlag && !*charsFlag {
		*bytesFlag = true
		*linesFlag = true
		*wordsFlag = true
	}
}

// getInputSource determines the input source based on arguments
func getInputSource(args []string) (io.ReadCloser, string) {
	if len(args) > 0 {
		filename := args[0]
		file, err := os.Open(filename)
		utils.CheckError(err, "Error opening file")
		return file, filename
	}
	return os.Stdin, ""
}

// collectResults gathers results from the counter based on active flags
func collectResults(c *counter.Counter, linesFlag, wordsFlag, bytesFlag, charsFlag bool) []int {
	var results []int
	if linesFlag {
		results = append(results, c.Lines())
	}
	if wordsFlag {
		results = append(results, c.Words())
	}
	if bytesFlag {
		results = append(results, c.Bytes())
	}
	if charsFlag {
		results = append(results, c.Characters())
	}
	return results
}
