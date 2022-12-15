package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const filename = "./input.txt"

func findTasksThatOverlapFully(content string) int {
	lines := strings.Split(content, "\n")
	count := 0
	for _, line := range lines {
		line := strings.Split(line, ",")
		fmt.Println(line)
		line1 := strings.Split(line[0], "-")
		line2 := strings.Split(line[1], "-")

		startA, _ := strconv.Atoi(line1[0])
		endA, _ := strconv.Atoi(line1[1])
		startB, _ := strconv.Atoi(line2[0])
		endB, _ := strconv.Atoi(line2[1])
		if (startB >= startA && endB <= endA) || (startA >= startB && endA <= endB) {
			count += 1
		}

	}
	return count
}

func findTasksThatOverlapAtAll(content string) int {
	lines := strings.Split(content, "\n")
	count := 0
	for _, line := range lines {
		line := strings.Split(line, ",")
		line1 := strings.Split(line[0], "-")
		line2 := strings.Split(line[1], "-")

		startA, _ := strconv.Atoi(line1[0])
		endA, _ := strconv.Atoi(line1[1])
		startB, _ := strconv.Atoi(line2[0])
		endB, _ := strconv.Atoi(line2[1])

	loop:
		for i := startA; i <= endA; i++ {
			for j := startB; j <= endB; j++ {
				if i == j {
					count++
					break loop
				}
			}
		}
	}
	return count
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
	fmt.Println(findTasksThatOverlapFully(content))
	fmt.Println(findTasksThatOverlapAtAll(content))

}
