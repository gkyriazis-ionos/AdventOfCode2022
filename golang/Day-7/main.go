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

func parseContent(i int, c string) []string {
	lines := strings.Split(c, "\n")
	input := strings.Split(lines[i], " ")
	return input
}
func workOnInput(c string) {
	var dir directory
	i := 0
	for {
		input := parseContent(i, c)
		if input[0] == "$" {
			if input[1] == "cd" {
				dir.location = input[2]
			} else if input[1] == "ls" {

				i++
				input = parseContent(i, c)

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
