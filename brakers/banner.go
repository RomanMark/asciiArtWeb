package brakers

import (
	"log"
	"os"
)

func ReadBanner(path string) []string {
	raw, err := os.ReadFile(path)
	var data []string
	if err != nil {
		log.Println(err)

		return data
	}
	var line string
	for _, char := range raw {
		if char == '\n' {
			if line == "" {
				continue
			}
			data = append(data, line)
			line = ""
			continue
		}
		line += string(char)
	}
	data = append(data, line)
	return data
}
