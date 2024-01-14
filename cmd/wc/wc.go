package wc

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type WordCount struct {
	file  string
	words int
	lines int
	bytes int
}

func wc(option string, file string) (string, error) {
	if option == "-c" {
		return getByteCount(file)
	} else if option == "-l" {
		return getLineCount(file)
	} else if option == "-w" {
		return getWordCount(file)
	}
	err := fmt.Sprintf("Invalid option")
	return "", errors.New(err)
}

func getByteCount(file string) (string, error) {
	// Read the file
	fd, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("%d %s", len(fd), file)
	return result, nil
}

func getLineCount(file string) (string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(fd)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	result := fmt.Sprintf("%d %s", lineCount, file)
	return result, nil
}

func getWordCount(file string) (string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
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
