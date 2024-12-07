package main

import (
	"os"
	"strconv"
	"strings"
)

type entry struct {
	result int64

	values []int64
}

func main() {
	input, err := os.ReadFile("7/input.txt")

	if err != nil {
		panic(err)
	}

	var entries []entry

	for _, lineRaw := range strings.Split(string(input), "\n") {
		line := strings.Split(lineRaw, ":")

		result, err := strconv.ParseInt(line[0], 10, strconv.IntSize)

		if err != nil {
			panic(err)
		}

		valuesRaw := strings.Split(strings.Trim(line[1], " "), " ")
		var values []int64
		for _, v := range valuesRaw {
			conv, err := strconv.ParseInt(v, 10, strconv.IntSize)
			if err != nil {
				panic(err)
			}
			values = append(values, conv)
		}

		entries = append(entries, entry{
			result: result,
			values: values,
		})
	}

	part1Process(entries)
	part2Process(entries)
}

func generateCombination(base string, ops string, maxLength int) []string {
	var combos []string

	if len(base) == maxLength {
		return []string{base}
	}

	for _, o := range ops {
		combos = append(combos, generateCombination(base+string(o), ops, maxLength)...)
	}

	return combos
}
