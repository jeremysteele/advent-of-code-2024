package main

import (
	"fmt"
	"math"
	"strconv"
)

func part1Process(rules map[string][]string, updates [][]string) {
	var middleTotal int64 = 0

	for _, update := range updates {
		valid := part1ValidateUpdate(rules, update)

		if valid {
			middle, err := strconv.ParseInt(update[int(math.Ceil(float64(len(update)/2)))], 10, strconv.IntSize)

			if err != nil {
				panic(err)
			}

			middleTotal += middle
		}

		fmt.Printf("%+v: %+v\n", update, valid)
	}

	fmt.Printf("Total: %d\n", middleTotal)
}

func part1ValidateUpdate(rules map[string][]string, update []string) bool {
	pastPages := make(map[string]bool)

	for _, page := range update {
		if pageRules, ok := rules[page]; ok {
			for _, rule := range pageRules {
				if _, ok := pastPages[rule]; ok {
					return false
				}
			}
		}

		pastPages[page] = true
	}

	return true
}
