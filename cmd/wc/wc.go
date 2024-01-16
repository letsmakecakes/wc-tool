package wc

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func WC(option string, file string) (string, error) {
	if option == "-c" {
		return getByteCount(file)
	} else if option == "-l" {
		return getLineCount(file)
	} else if option == "-w" {
		return getWordCount(file)
	} else if option == "-m" {
		return getCharacterCount(file)
	} else if option == "" {
		return getAllCount(file)
	}
	err := fmt.Sprintf("Invalid option")
	return "", errors.New(err)
}

// Function to get the byte count in a file
func getByteCount(file string) (string, error) {
	// Read the file
	fd, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("%d %s", len(fd), file)
	return result, nil
}

// Function to get the line count in a file
func getLineCount(file string) (string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {

		}
	}(fd)
	scanner := bufio.NewScanner(fd)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	result := fmt.Sprintf("%d %s", lineCount, file)
	return result, nil
}

// Function to get word count in a file
func getWordCount(file string) (string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {

		}
	}(fd)

	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanWords)

	countWords := 0
	for scanner.Scan() {
		countWords++
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	result := fmt.Sprintf("%d %s", countWords, file)
	return result, nil
}

// Function to get character count
func getCharacterCount(file string) (string, error) {
	// Open the file
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer func(fd *os.File) {
		err := fd.Close()
		if err != nil {

		}
	}(fd)

	// Create a new bufio reader
	reader := bufio.NewReader(fd)

	// Create a variable to store the number of characters
	var numChars int

	// Read the file line by line
	for {
		// Read a character from the file
		_, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		numChars++
	}

	// return the number of characters
	result := fmt.Sprintf("%d %s", numChars, file)
	return result, nil
}

// Function get count of bytes, lines and words
func getAllCount(file string) (string, error) {
	byteCount, err := getByteCount(file)
	if err != nil {
		return "", err
	}
	lineCount, err := getLineCount(file)
	if err != nil {
		return "", err
	}
	wordCount, err := getWordCount(file)
	if err != nil {
		return "", err
	}

	bytes := strings.Split(byteCount, " ")
	lines := strings.Split(lineCount, " ")
	words := strings.Split(wordCount, " ")

	result := fmt.Sprintf("%s   %s  %s %s", lines[0], words[0], bytes[0], file)
	return result, nil
}
