package wc

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

func WC(option string, file string) (string, error) {
	if isByteCount(option) {
		return getByteCount(file)
	}

	if isLineCount(option) {
		return getLineCount(file)
	}

	if isWordCount(option) {
		return getWordCount(file)
	}

	if isCharacterCount(option) {
		return getCharacterCount(file)
	}

	if isEmpty(option) {
		return getAllCount(file)
	}

	err := fmt.Sprintf("Invalid option")
	return "", errors.New(err)
}

func isByteCount(option string) bool {
	return option == "-c"
}

func isLineCount(option string) bool {
	return option == "-l"
}

func isWordCount(option string) bool {
	return option == "-w"
}

func isCharacterCount(option string) bool {
	return option == "-m"
}

func isEmpty(option string) bool {
	return option == ""
}

// Function to get the byte count in a file
func getByteCount(file string) (string, error) {
	fd := openFile(file)
	defer closeFile(fd)

	bytes, err := countBytes(fd)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%d %s", bytes, file)
	return result, nil
}

func countBytes(fd *os.File) (int, error) {
	reader := bufio.NewReader(fd)
	byteCount := 0

	for {
		_, err := reader.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
		byteCount++
	}

	return byteCount, nil
}

// Function to get the line count in a file
func getLineCount(file string) (string, error) {
	fd := openFile(file)
	defer closeFile(fd)

	lines, err := countLines(fd)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%d %s", lines, file)
	return result, nil
}

func countLines(fd *os.File) (int, error) {
	scanner := bufio.NewScanner(fd)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount, nil
}

// Function to get word count in a file
func getWordCount(file string) (string, error) {
	fd := openFile(file)
	defer closeFile(fd)

	words, err := countWords(fd)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%d %s", words, file)
	return result, nil
}

func countWords(fd *os.File) (int, error) {
	scanner := bufio.NewScanner(fd)
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}

// Function to get character count
func getCharacterCount(file string) (string, error) {
	fd := openFile(file)
	defer closeFile(fd)

	characters, err := countCharacters(fd)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%d %s", characters, file)
	return result, nil
}

func countCharacters(fd *os.File) (int, error) {
	reader := bufio.NewReader(fd)
	var characterCount int

	for {
		_, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		characterCount++
	}

	return characterCount, nil
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

func openFile(file string) *os.File {
	fd, err := os.Open(file)
	if err != nil {
		panic("error in opening the file")
	}

	return fd
}

func closeFile(file *os.File) {
	func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Errorf("error in closing the file: %v", err)
		}
	}(file)
}
