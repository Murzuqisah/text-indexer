package parser

import (
	"fmt"
	"os"
)

// mmap(memory-mapped files) reads large files efficiently
// read the file form the argument
func Reader(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	return file
}
