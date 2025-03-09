package parser

import (
	"bufio"
	"slices"
)

type Chunk struct {
	Offset int64
	Data   []byte
}

// chunks the files into fixed-size chunks
func ChunkFile(filePath string, chunkSize int) ([]Chunk, error) {
	var chunks []Chunk

	scanner := bufio.NewScanner(Reader(filePath))
	buf := make([]byte, chunkSize)
	scanner.Buffer(buf, chunkSize)

	offset := int64(0)
	for scanner.Scan() {
		chunks = append(chunks, Chunk{Offset: offset, Data: slices.Clone(scanner.Bytes())})
		offset += int64(len(scanner.Bytes()))
	}
	return chunks, scanner.Err()
}
