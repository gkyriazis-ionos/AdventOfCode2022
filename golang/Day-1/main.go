package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filename = "./input.txt"

func getContent(location string) (string, error) {
	content, err := os.ReadFile(location)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func findMaxBlock(content string) (int, error) {
	lines := strings.SplitAfter(content, "\n")

	max := 0
	sum := 0

	for _, line := range lines {
		line = strings.Trim(line, "\n")
		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			return -1, err
		}
		sum += number
	}
	return sum, nil
}

func findSumOfMaxBlock(content string) (int, error) {
	lines := strings.SplitAfter(content, "\n")

	sum := 0
	var sums []int

	for _, line := range lines {
		line = strings.Trim(line, "\n")
		if line == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			return -1, err
		}
		sum += number
	}
	sums = append(sums, sum)

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	return sums[0] + sums[1] + sums[2], nil
}

func main() {
	content, err := getContent(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	_, err = findMaxBlock(content)
	if err != nil {
		log.Fatalf("%v", err)
	}

	sumMax, err := findSumOfMaxBlock(content)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("Max %d\n", sumMax)
}
