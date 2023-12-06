package env

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Getenv(key string) string {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var value string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		segments := strings.Split(line, "=")

		if len(segments) == 2 && segments[0] == key {
			value = strings.ReplaceAll(segments[1], `"`, "")
		}
	}

	return value
}