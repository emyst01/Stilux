package read

import (
	"bufio"
	"fmt"
	"os"
)

type CustomError struct {
	errType string
	Content string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.errType, e.Content)
}

func Read() (string, error) {
	var fileContent string
	file, err := os.Open("index.stilux")
	if err != nil {
		return "", CustomError{
			errType: "FileOpenError",
			Content: "There was an error reading the file:\n" + err.Error(),
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fileContent += line
	}

	if err := scanner.Err(); err != nil {
		return "", CustomError{
			errType: "ScannerError",
			Content: "Error scanning the file:\n" + err.Error(),
		}
	}

	return fileContent, nil
}
