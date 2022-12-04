package advent

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(lines chan string) {
	fileContent, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fileContent)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		lines <- scanner.Text()
	}
	close(lines)
}
