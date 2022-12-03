// Main package for rock paper Scissors
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Player struct
type Player struct {
	score  int
	choice string
}

func main() {

	// legend := map[string]string{
	// 	"A": "Rock",
	// 	"B": "Paper",
	// 	"C": "Scissors",
	// 	"X": "Rock",
	// 	"Y": "Paper",
	// 	"Z": "Scissors",
	// }

	win := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
		"X": "C",
		"Y": "A",
		"Z": "B",
	}

	lose := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
		"X": "B",
		"Y": "C",
		"Z": "A",
	}

	draw := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	points := map[string]int{
		"A":    1,
		"B":    2,
		"C":    3,
		"X":    1,
		"Y":    2,
		"Z":    3,
		"win":  6,
		"lose": 0,
		"draw": 3,
	}

	file, err := os.Open("rps.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	allWinTotal := 0

	planTotal := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Split(line, " ")[1] == "" {
			continue
		}

		opp := strings.Split(line, " ")[0]
		me := strings.Split(line, " ")[1]

		fmt.Println("Opponent: ", opp)
		fmt.Println("Me: ", me)
		fmt.Println(win[opp], lose[opp], draw[opp])
		fmt.Println(win[me], lose[me], draw[me])
		if win[me] == opp {

			score := points["win"]
			score += points[me]
			fmt.Println(score)
			allWinTotal += score
		} else if lose[me] == opp {
			score := points["lose"]
			score += points[me]
			fmt.Println(score)
			allWinTotal += score
		} else if draw[me] == opp {
			score := points["draw"]
			score += points[me]
			fmt.Println(score)
			allWinTotal += score
		}

		if me == "X" {
			score := points[win[opp]]
			score += points["lose"]
			planTotal += score
		} else if me == "Y" {
			score := points[opp]
			score += points["draw"]
			planTotal += score
		} else if me == "Z" {
			score := points[lose[opp]]
			score += points["win"]
			planTotal += score
		}

	}
	fmt.Printf("All win: %d\n", allWinTotal)
	fmt.Printf("Plan: %d\n", planTotal)
}
