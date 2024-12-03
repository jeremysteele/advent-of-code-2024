package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// This is some garbage, but it works.
func main() {
	file, err := os.ReadFile("3/input.txt")

	// open file
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)

	var re = regexp.MustCompile(`(?m)(mul\([0-9]+,[0-9]+\))?(do\(\))?(don\'t\(\))?`)

	var total int64 = 0
	var do = true

	for _, match := range re.FindAllStringSubmatch(input, -1) {
		for i, submatch := range match {
			if i == 0 {
				continue
			}

			if submatch == "" {
				continue
			}

			if submatch[0:3] == "don" {
				do = false
			} else if submatch[0:2] == "do" {
				do = true
			} else {
				split := strings.Split(submatch[4:len(submatch)-1], ",")

				first, err := strconv.ParseInt(split[0], 10, strconv.IntSize)
				if err != nil {
					panic(err)
				}

				second, err := strconv.ParseInt(split[1], 10, strconv.IntSize)
				if err != nil {
					panic(err)
				}

				fmt.Printf("Multiplying %d, %d\n", first, second)

				if do {
					total += first * second
				}
			}

			fmt.Printf("%+v\n", submatch)
		}
	}

	fmt.Printf("Total: %d\n", total)
}
