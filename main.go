package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Automata interface
type Automata interface {
	isValid(string) (bool, string)
}

// DFA implementation deterministic finite automaton
type dfa struct {
	alphabet string
	start    int
	success  []int
	matrix   [][]int
}

// isValid implementation
func (a dfa) isValid(word string) (bool, string) {
	current := a.start
	state := strconv.Itoa(a.start)

	for _, item := range word {
		current = a.matrix[current][strings.Index(a.alphabet, string(item))]
		state = fmt.Sprintf("%s/%d", state, current)
	}

	return contains(a.success, current), state
}

func validate(a Automata, word string) {
	isValid, state := a.isValid(word)
	if isValid {
		fmt.Println("ACEPTADO")
	} else {
		fmt.Println("RECHAZADO")
	}

	fmt.Println(state)
}

func main() {
	var automata Automata
	// Filename input
	name := flag.String("file", "in.txt", "File name")
	flag.Parse()

	// Pass the filename, opens the file and gets a scanner
	file, err := os.Open(*name)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read string
	word := scanLine(scanner)

	// Read alphabet
	alphabet := scanAlphabet(scanner)

	// Read start
	start, _ := strconv.Atoi(scanLine(scanner))

	// Read final states
	success := scanFinalStates(scanner)

	// Read function
	matrix := scanMatrix(scanner)

	// Create DFA
	automata = &dfa{
		alphabet,
		start,
		success,
		matrix,
	}

	// Validate a given word
	validate(automata, word)
}

func scanLine(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}

func scanAlphabet(scanner *bufio.Scanner) string {
	line := scanLine(scanner)

	alphabet := ""
	for _, item := range strings.Split(line, ";") {
		if len(item) < 1 {
			continue
		}
		alphabet = fmt.Sprintf("%s%s", alphabet, item)
	}

	return alphabet
}

func scanFinalStates(scanner *bufio.Scanner) []int {
	line := scanLine(scanner)
	success := []int{}
	for _, item := range strings.Split(line, ";") {
		number, err := strconv.Atoi(item)
		if err != nil {
			continue
		}
		success = append(success, number)
	}

	return success
}

func scanMatrix(scanner *bufio.Scanner) [][]int {
	matrix := [][]int{}
	for scanner.Scan() {
		array := []int{}
		for _, item := range strings.Split(scanner.Text(), ";") {
			number, err := strconv.Atoi(item)
			if err != nil {
				continue
			}
			array = append(array, number)
		}
		matrix = append(matrix, array)
	}

	return matrix
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
