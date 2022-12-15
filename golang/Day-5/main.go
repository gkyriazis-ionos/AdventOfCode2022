package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

const filename = "./input.txt"

func getContent(location string) (string, error) {
	content, err := os.ReadFile(location)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

type parseInterface interface {
	parseCrates(s string) [][]string
	parseMoves(s string) [][]int
}

type input struct {
	scanner string
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func (in input) parseCrates(s string) []string {
	lines := strings.Split(s, "\n")
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
	//crates2 := make([][]string, len(crates))
	crates2 := crates[:len(crates)-1]

	//copy(crates2, crates)
	fmt.Println(crates2)
	var st []string
	for j := range crates2 {
		//for j := 0; j < len(crates); j++ {
		sb := strings.Builder{}
		for i := range crates2[j] {
			//fmt.Println(crates2[i][j])
			//for i := 0; i < len(crates[j]); i++ {
			_, _ = sb.WriteString(crates2[i][j])
		}
		revSt := reverse(sb.String())
		revSt = strings.TrimSpace(revSt)
		st = append(st, revSt)
	}
	fmt.Println(st)

	return st
}

func (in input) parseMoves(s string) [][]int {
	lines := strings.Split(s, "\n")
	var rows []string
	found := false
	for _, line := range lines {
		if len([]rune(line)) == 0 {
			found = true
		}
		if len([]rune(line)) != 0 && !found {
			continue
		}
		rows = append(rows, line)
	}
	var commands [][]string
	for _, row := range rows {
		command := strings.Split(row, "\n")
		commands = append(commands, command)
	}
	var moves [][]string
	//fmt.Println(commands)
	_, commands = commands[0], commands[1:]
	//fmt.Println(commands)
	re := regexp.MustCompile("[0-9]+")
	for i := 0; i < len(commands); i++ {
		for j := 0; j < len(commands[i]); j++ {
			moves = append(moves, re.FindAllString(commands[i][j], -1))
			//fmt.Println(moves)
		}
	}

	fmt.Println(moves[1][1])
	var moves2 [][]int

	for i := 0; i < len(moves); i++ {
		moves2 = append(moves2, []int{})

		for j := 0; j < len(moves[i]); j++ {
			dd, _ := strconv.Atoi(moves[i][j])
			moves2[i] = append(moves2[i], dd)
		}
	}

	return moves2
}

func calcOutputPt1(crates []string, moves [][]int) string {
	for i := 0; i < len(moves); i++ {
		for j := 1; j < len(moves[i]); j++ {
			moves[i][j] = moves[i][j] - 1
		}
	}

	for _, move := range moves {
		num := move[0]
		from := move[1]
		to := move[2]
		substr := crates[from][len(crates[from])-num : len(crates[from])]
		crates[to] = crates[to] + reverse(substr)
		crates[from] = crates[from][0 : len(crates[from])-num]
	}

	fmt.Println(crates)
	output := ""
	for _, st := range crates {
		output = output + string(st[len(st)-1])
	}

	return output
}

func calcOutputPt2(crates []string, moves [][]int) string {
	for _, move := range moves {
		fmt.Println(move)
		num := move[0]
		from := move[1]
		to := move[2]
		substr := crates[from][len(crates[from])-num : len(crates[from])]
		crates[to] = crates[to] + substr
		crates[from] = crates[from][0 : len(crates[from])-num]
	}

	fmt.Println(crates)
	output := ""
	for _, st := range crates {
		output = output + string(st[len(st)-1])
	}

	return output
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
	crates := s.parseCrates(s.scanner)
	fmt.Println(crates)
	moves := s.parseMoves(s.scanner)
	//fmt.Println(moves)
	finalAnswerPt1 := calcOutputPt1(crates, moves)
	fmt.Println(finalAnswerPt1)
	crates = s.parseCrates(s.scanner)
	//fmt.Println(crates)
	finalAnswerPt2 := calcOutputPt2(crates, moves)
	fmt.Println(finalAnswerPt2)

}
