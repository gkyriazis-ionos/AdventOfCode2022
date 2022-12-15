package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const filename = "./input.txt"

func calcScore(c string) int {
	content, err := getContent(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	score := 0
	lines := strings.Split(content, "\n")

	for _, line := range lines {

		a, b, _ := strings.Cut(line, " ")
		switch a {
		case "A":
			switch b {
			case "X":
				score += 3 + 1
			case "Y":
				score += 6 + 2
			case "Z":
				score += 0 + 3
			}
		case "B":
			switch b {
			case "X":
				score += 0 + 1
			case "Y":
				score += 3 + 2
			case "Z":
				score += 6 + 3
			}
		case "C":
			switch b {
			case "X":
				score += 6 + 1
			case "Y":
				score += 0 + 2
			case "Z":
				score += 3 + 3
			}

		}

	}
	return score
}

func calcScorePart2(c string) int {
	content, err := getContent(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	score := 0
	lines := strings.Split(content, "\n")

	for _, line := range lines {

		a, b, _ := strings.Cut(line, " ")
		switch a {
		case "A":
			switch b {
			case "X":
				score += 0 + 3
			case "Y":
				score += 3 + 1
			case "Z":
				score += 6 + 2
			}
		case "B":
			switch b {
			case "X":
				score += 0 + 1
			case "Y":
				score += 3 + 2
			case "Z":
				score += 6 + 3
			}
		case "C":
			switch b {
			case "X":
				score += 0 + 2
			case "Y":
				score += 3 + 3
			case "Z":
				score += 6 + 1
			}

		}

	}
	return score
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

	score := calcScore(content)
	fmt.Printf("Score PT 1: %d\n", score)

	score = calcScorePart2(content)
	fmt.Printf("Score PT 2: %d\n", score)
}
