package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
)

const filename = "./example-input.txt"

func getContent(location string) (string, error) {
	content, err := os.ReadFile(location)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

type parseInterface interface {
	parseCrates(s string) [][]string
	parseMoves(s string)
}

type input struct {
	scanner string
}

func (in input) parseCrates(s string) [][]string {
	lines := strings.Split(s, "\n")
	//blankIndex := -1
	var rows []string

	for _, line := range lines {
		if len([]rune(line)) == 0 {
			break
		}
		rows = append(rows, line)
	}
	var crates [][]string
	for i := 0; i < len(rows); i++ {
		crates = append(crates, []string{})
		byteArray, err := syscall.ByteSliceFromString(rows[i])
		if err != nil {
			fmt.Println("Error in the byte array")
		}
		for j := 1; j < len(byteArray); j += 4 {
			crates[i] = append(crates[i], string(byteArray[j]))
		}
	}
	maxLen := -1
	for i := 0; i < len(crates); i++ {
		if len(crates[i]) > maxLen {
			maxLen = len(crates[i])
		}
	}
 	for i := 0; i < len(crates); i++ {
		 for len(crates[i]) < maxLen {
			 crates[i] = append(crates[i], " ")
		 }
	}
	fmt.Println(crates)
	return [][]string{}
}

func main() {
	content, err := getContent(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	//fmt.Println(content)
	s := input{
		scanner: content,
	}
	s.parseCrates(s.scanner)
}
