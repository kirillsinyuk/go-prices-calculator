package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("Error opening file: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.New("Error scanning file: " + err.Error())
	}
	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Error opening file: " + err.Error())
	}
	defer file.Close()

	time.Sleep(3 * time.Second) // delay to simulate slow operation

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Error encoding file: " + err.Error())
	}
	return nil
}

func New(inputFilePath, outputFilePath string) *FileManager {
	return &FileManager{inputFilePath, outputFilePath}
}
