package reader

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var wordFile = "dictionary/wordninja_words.txt"

// loadWords loads the engish split words from file, return list of words.
func LoadWords() []string {
	words, err := readFileByLine()
	if err != nil {
		panic("load english split word failed, " + err.Error())
	}

	return words
}

// readFileByLine returns a list by reading file line by line
func readFileByLine() ([]string, error) {
	var lines []string
	f, err := os.Open(wordFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if line == "\n" {
			continue
		}

		line = strings.Replace(line, "\n", "", -1)
		lines = append(lines, line)
	}

	return lines, nil
}
