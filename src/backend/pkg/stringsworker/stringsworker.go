package stringsworker

import (
	"bufio"
	"log"
	"os"
)

func ReadToken(path string) (token string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return token
}
