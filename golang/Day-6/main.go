package main

import (
	"fmt"
	"log"
	"os"
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

func parseInputAndFindNumOfChars(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		charsArray := strings.Split(line, "")
		fmt.Println(charsArray)
		for i := 0; i < len(charsArray); i++ {
			if charsArray[i] != charsArray[i+1] && charsArray[i] != charsArray[i+2] && charsArray[i] != charsArray[i+3] &&
				charsArray[i+1] != charsArray[i+2] && charsArray[i+1] != charsArray[i+3] && charsArray[i+2] != charsArray[i+3] {
				fmt.Println(i + 3 + 1)
				break
			}
		}

	}
}

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}

func parseInputAndFindNumOfMsg(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		charsArray := strings.Split(line, "")
		var found []string
		for i := 0; i < len(charsArray); i++ {
			found = append(found, charsArray[i])
			//fmt.Println(found)
			if i > 14 {
				s := strings.Join(found[i-14:i], "")
				if unique(s) {
					fmt.Println(i)
					break
				}
			}
		}
	}
}

func main() {
	content, err := getContent(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	//fmt.Println(content)
	//parseInputAndFindNumOfChars(content)
	parseInputAndFindNumOfMsg(content)

}
