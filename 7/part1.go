package main

import (
	"fmt"
)

const part1Operators = "+*"

/*func debugOutput(data []int64, ops string) string {
	output := ""

	for i, num := range data {
		if i == 0 {
			output += fmt.Sprintf("%s %d", strings.Repeat("(", len(ops)), num)
		} else {
			output += fmt.Sprintf("%c %d ) ", ops[i-1], num)
		}
	}

	return output
}*/

func part1ProcessLine(result int64, data []int64) bool {
	operatorCombos := generateCombination("", part1Operators, len(data)-1)

	for _, combo := range operatorCombos {
		total := int64(0)

		for i, num := range data {
			if i == 0 {
				total = num
				continue
			}

			switch combo[i-1] {
			case '+':
				total += num
			case '*':
				total *= num
			}
		}

		if total == result {
			//fmt.Printf("%d == %s\n", result, debugOutput(data, combo))
			return true
		}
	}

	return false
}

func part1Process(entries []entry) {
	total := int64(0)

	for _, e := range entries {
		if part1ProcessLine(e.result, e.values) {
			total += e.result
		}
	}

	fmt.Printf("Total: %d\n", total)
}
