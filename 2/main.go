package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func cloneWithoutIndex(line []int, index int) []int {
	newSlice := make([]int, 0)
	newSlice = append(newSlice, line[:index]...)
	newSlice = append(newSlice, line[index+1:]...)
	return newSlice
}

func check(line []int) bool {
	prev := -5
	next := -5
	increasing := false

	for i, level := range line {
		next = level

		if i == 0 {
			prev = next
			continue
		}

		diff := math.Abs(float64(prev - next))

		if diff < 1 || diff > 3 {
			return false
		}

		if i == 1 {
			increasing = prev < next
		}

		if increasing != (prev < next) {
			return false
		}

		prev = next
	}

	return true
}

// This is some garbage, but it works.
func main() {
	// open file
	file, err := os.Open("2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numSafe := 0
	for scanner.Scan() {
		splitStr := strings.Split(scanner.Text(), " ")

		line := make([]int, len(splitStr))

		for i := range splitStr {
			readNext, err := strconv.ParseInt(splitStr[i], 10, strconv.IntSize)
			if err != nil {
				panic(err)
			}

			line[i] = int(readNext)
		}

		if check(line) {
			numSafe++
			continue
		}

		fmt.Printf("Unsafe: %+v\n", splitStr)

		// Part 2
		for i := range line {
			if check(cloneWithoutIndex(line, i)) {
				numSafe++
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Num safe: %d\n", numSafe)

}
