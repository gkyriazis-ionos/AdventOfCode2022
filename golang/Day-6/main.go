package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const filename = "./example-input.txt"

func getContent(location string) (os.File, error) {
	content, err := os.Open(location)
	if err != nil {
		return os.File{}, err
	}
	return *content, nil
}

func readInput(input bufio.Scanner) [][]string {
	var patchOfBoxes [][]string
	rowCount := 0
	for input.Scan() {
		lineScanner := bufio.NewScanner(strings.NewReader(input.Text()))
		lineScanner.Split(bufio.ScanWords)
		for lineScanner.Scan() {
			if len(patchOfBoxes) <= rowCount {
				patchOfBoxes = append(patchOfBoxes, []string{})
			}
			var height string
			if _, err := fmt.Sscan(lineScanner.Text(), &height); err != nil {
				log.Fatal(err)
			}
			if strings.Join(patchOfBoxes[rowCount], "") == " " {
				break
			}
			patchOfBoxes[rowCount] = append(patchOfBoxes[rowCount], height)
			fmt.Println(patchOfBoxes[rowCount])

		}
		rowCount += 1
	}
	return patchOfBoxes

}

func main() {
	content, err := getContent(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	//fmt.Println(content)
	scanner := bufio.NewScanner(&content)
	boxes := readInput(*scanner)
	fmt.Println(boxes)
}
