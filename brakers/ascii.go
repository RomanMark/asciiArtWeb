package brakers

import (
	"strings"
)

func PrintAscii(text string, style string) string {
	path := "banners/standart.txt"
	switch style {
	case "standard":
		path = "banners/standard.txt"
	case "thinkertoy":
		path = "banners/thinkertoy.txt"
	case "shadow":
		path = "banners/shadow.txt"
	}

	data := ReadBanner(path)
	var out string
	if len(data) == 0 {
		return out
	}

	words := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	for i, word := range words {
		if i > 0 && word == "" {
			out += "\n"
		}
		if word == "" {
			continue
		}
		for j := 0; j < 8; j++ {
			for _, char := range []byte(word) {
				if int(char) < 32 || int(char) > 126 {
					return ""
				}
				out += data[int(char-32)*8+j]
			}
			out += "\n"
		}
	}
	return out
}
