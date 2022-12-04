// Main package
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	priority = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
)

// Rucksack is a struct that holds the items
type Rucksack struct {
	items []string
}

func main() {
	rucksack, err := Load()
	if err != nil {
		log.Fatal(err)
	}

	rucksack.SolveProblem1()
	rucksack.SolveProblem2()

}

// SolveProblem1 solves the first problem
func (rucksack *Rucksack) SolveProblem1() {
	// Solve the problem
	var prio int
	prio = 0
	for _, item := range rucksack.items {
		half := len(item) / 2
		for _, char := range item[:half] {
			if strings.ContainsRune(item[half:], char) {
				prio += priority[string(char)]
				break
			}
		}
	}
	fmt.Println(prio)
}

// SolveProblem2 solves the second problem
func (rucksack *Rucksack) SolveProblem2() {
	// Solve the problem
	var prio int
	prio = 0
	for i := 0; i < len(rucksack.items); i += 3 {
		for _, item := range rucksack.items[i] {
			if strings.ContainsRune(rucksack.items[i+1], item) && strings.ContainsRune(rucksack.items[i+2], item) {
				prio += priority[string(item)]
				break
			}
		}
	}

	fmt.Println(prio)
}

// Load loads the items from the file
func Load() (*Rucksack, error) {
	// Load the rucksack
	f, err := os.Open("rucksack.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rucksack := Rucksack{
		items: make([]string, 0),
	}

	// Read the rucksack
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		rucksack.items = append(rucksack.items, line)
	}

	return &rucksack, nil
}
