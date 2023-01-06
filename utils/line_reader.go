package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type LineReader struct {
	Line    int
	file    *os.File
	scanner *bufio.Scanner
}

func NewLineReader(fileName string) *LineReader {
	log.Printf("Opening %v...", fileName)
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return &LineReader{
		file:    file,
		scanner: bufio.NewScanner(file),
	}
}

func (lr *LineReader) Close() error {
	return lr.file.Close()
}

func (lr *LineReader) HasNext() bool {
	result := lr.scanner.Scan()

	if result {
		lr.Line++
	} else {
		Close(lr.file)
	}

	return result
}

func (lr *LineReader) Text() string {
	return lr.scanner.Text()
}

func (lr *LineReader) Integers() []int {
	as := strings.Split(lr.Text(), ",")
	result := make([]int, len(as))

	for i, s := range as {
		result[i] = Must(strconv.Atoi(s))
	}

	return result
}
