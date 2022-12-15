package main

import (
	"log"
	"os"
	"strings"
)

const filename = "./input.txt"

type files struct {
	size int
	name string
}

type directory struct {
	location       string
	subdirectories []string
	files          files
}

func parseContent(c string) {
	var dir directory
	lines := strings.Split(c, "\n")
	for _, line := range lines {
		input := strings.Split(line, " ")

		if input[0] == "$" {
			if input[1] == "cd" {
				dir.location = input[2]
			} else if input[1] == "ls" {
				
			}

		}

	}
}

func getContent(location string) (string, error) {
	content, err := os.ReadFile(location)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	content, err := getContent(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	parseContent(content)
}
