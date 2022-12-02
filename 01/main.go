// Main package
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Elves struct
type Elves struct {
	elves []int64
}

func createElfs() (*Elves, error) {
	// Open the file
	file, err := os.Open("calories.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	elves := Elves{
		elves: make([]int64, 0),
	}

	total := int64(0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves.elves = append(elves.elves, total)
			total = 0
		} else {
			var calories int64
			calories, err = strconv.ParseInt(line, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			total += calories
		}
		fileLines = append(fileLines, scanner.Text())
	}

	return &elves, nil
}

// SortInt64 sorts a slice of int64 in ascending order
func SortInt64(slice []int64) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// SortInt64Desc sorts a slice of int64 in descending order
func SortInt64Desc(slice []int64) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[j] < slice[i]
	})
}

func main() {
	elves, err := createElfs()
	if err != nil {
		log.Fatal(err)
	}

	SortInt64Desc(elves.elves)

	fmt.Println(elves.elves[0])

	top := elves.elves[0] + elves.elves[1] + elves.elves[2]

	fmt.Println(top)

}
