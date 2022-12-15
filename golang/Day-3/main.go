package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const filename = "./input.txt"

func findPrio() map[string]int {
	prio := map[string]int{}
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	chars := strings.Split(alphabet, "")
	for i, c := range chars {
		prio[c] = i + 1
	}
	return prio

}

func findBadgePrio(content string) int {
	lines := strings.Split(content, "\n")

	var found []string
	for i := 0; i < len(lines); i += 3 {
		elf1 := lines[i]
		elf2 := lines[i+1]
		elf3 := lines[i+2]
		for _, c := range elf1 {
			if strings.Contains(elf2, string(c)) && strings.Contains(elf3, string(c)) {
				found = append(found, string(c))
				break
			}
		}
	}
	//fmt.Println(found)
	prio := findPrio()
	sum := 0
	for _, c := range found {
		sum += prio[c]
	}

	return sum
}

func findItemsPrio(content string) int {
	lines := strings.Split(content, "\n")

	var found []string
	for _, line := range lines {
		chars := strings.Split(line, "")
		bag1 := chars[0 : len(chars)/2]
		bag2 := chars[len(chars)/2:]
		for _, item := range bag1 {
			if strings.Contains(strings.Join(bag2, ""), item) {
				found = append(found, item)
				break
			}
		}
	}
	prio := findPrio()
	sum := 0
	for _, c := range found {
		sum += prio[c]
	}

	return sum
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
	fmt.Println(findItemsPrio(content))
	fmt.Println(findBadgePrio(content))
}
