package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const filename = "./input.txt"

func getContent(location string) (os.File, error) {
	content, err := os.Open(location)
	if err != nil {
		return os.File{}, err
	}
	return *content, nil
}

func readInput(input bufio.Scanner) [][]uint8 {
	var patchOfTrees [][]uint8
	rowCount := 0
	for input.Scan() {
		lineScanner := bufio.NewScanner(strings.NewReader(input.Text()))
		lineScanner.Split(bufio.ScanRunes)
		for lineScanner.Scan() {
			if len(patchOfTrees) <= rowCount {
				patchOfTrees = append(patchOfTrees, []uint8{})
			}
			var height uint8
			if _, err := fmt.Sscan(lineScanner.Text(), &height); err != nil {
				log.Fatal(err)
			}
			patchOfTrees[rowCount] = append(patchOfTrees[rowCount], height)
		}
		rowCount += 1
	}
	return patchOfTrees
}

func countOfTheOuterTrees(input [][]uint8) int {
	return 2*len(input) + 2*len(input[0]) - 4
}

func isTreeVisible(x int, y int, input [][]uint8) bool {
	//count := 0
	visible := true
	for i := 0; i < x; i++ {
		if input[i][y] >= input[x][y] {
			visible = false
		}
	}
	if visible {
		return true
	}
	visible = true
	for i := x + 1; i < len(input); i++ {
		if input[i][y] >= input[x][y] {
			visible = false
		}
	}
	if visible {
		return true
	}
	visible = true
	for j := 0; j < y; j++ {
		if input[x][j] >= input[x][y] {
			visible = false
		}
	}
	if visible {
		return true
	}
	visible = true
	for j := y + 1; j < len(input[x]); j++ {
		if input[x][j] >= input[x][y] {
			visible = false
		}
	}
	if visible {
		return true
	}
	return false
}

func countViewScore(x int, y int, input [][]uint8) int {
	// left
	leftScore := 0
	for i := x - 1; i >= 0; i-- {
		if input[i][y] < input[x][y] {
			leftScore++
		}
		if input[i][y] >= input[x][y] {
			leftScore++
			break
		}
	}
	// right
	rightScore := 0
	for i := x + 1; i < len(input); i++ {
		if input[i][y] < input[x][y] {
			rightScore++
		}
		if input[i][y] >= input[x][y] {
			rightScore++
			break
		}
	}

	// top
	topScore := 0
	for j := y - 1; j >= 0; j-- {
		if input[x][j] < input[x][y] {
			topScore++
		}
		if input[x][j] >= input[x][y] {
			topScore++
			break
		}
	}
	// right
	botScore := 0
	for j := y + 1; j < len(input[x]); j++ {
		if input[x][j] < input[x][y] {
			botScore++
		}
		if input[x][j] >= input[x][y] {
			botScore++
			break
		}
	}

	return leftScore * rightScore * topScore * botScore

}

func main() {
	content, err := getContent(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	scanner := bufio.NewScanner(&content)
	input := readInput(*scanner)

	count := 0
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if isTreeVisible(i, j, input) {
				count++
			}
		}
	}
	fmt.Println(count + countOfTheOuterTrees(input))

	max := -1
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if countViewScore(i, j, input) > max {
				max = countViewScore(i, j, input)
			}
		}
	}
	fmt.Println(max)

}
