package utils

import (
	"bufio"
	"log"
	"os"
)

func StreamFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var contents []byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, []byte(scanner.Text())...)
	}

	return contents
}
