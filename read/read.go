package read

import (
	"bufio"
	"os"
)

func ReadExactLine(fileName string, lineNumber int) (string, error) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	br := bufio.NewReader(inputFile)
	for i := 1; i < lineNumber; i++ {
		_, _ = br.ReadString('\n')
	}
	str, err := br.ReadString('\n')
	if err != nil {
		return "", err
	}
	return str, nil
}
