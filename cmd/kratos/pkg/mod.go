package pkg

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const modFile = "go.mod"

func GetModName() string {
	fileName := modFile
	for {
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fileName = "../" + fileName
		} else {
			break
		}
	}
	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return ""
	}

	buf := bufio.NewReader(f)
	line, err := buf.ReadString('\n')
	if err != nil {
		log.Println(err)
		return ""
	}
	log.Println(string(line))
	line = strings.TrimSpace(line)
	modLines := strings.Split(line, " ")
	if len(modLines) != 2 {
		return ""
	}
	if modLines[0] == "module" {
		return modLines[1]
	}
	return ""
}
