package internal

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(loc string) []string {
	file, err := os.Open(loc)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err != nil {
		log.Fatal(err)
	}

	return lines
}

func Reader(loc string) (*bufio.Reader, *os.File) {
	file, err := os.Open(loc)

	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewReader(file), file
}
