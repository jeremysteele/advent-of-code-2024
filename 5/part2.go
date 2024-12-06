package main

import (
	"fmt"
	"math"
	"strconv"
)

const maxRetries = 200

func part2Process(rules map[string][]string, updates [][]string) {
	var middleTotal int64 = 0

	for i, update := range updates {
		valid := part2ProcessUpdate(rules, i, update, 0)

		if !valid {
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

func part2ProcessUpdate(rules map[string][]string, updateIndex int, update []string, retryCount int) bool {
	if retryCount > maxRetries {
		panic("Max retries hit")
	}

	pastPages := make(map[string]int)
	validUpdate := []string{}

	for i, page := range update {
		if pageRules, ok := rules[page]; ok {
			for _, rule := range pageRules {
				if pastIndex, ok := pastPages[rule]; ok {
					update[i] = update[pastIndex]
					update[pastIndex] = page
					part2ProcessUpdate(rules, updateIndex, update, retryCount+1)
					return false
				}
			}
		}

		validUpdate = append(validUpdate, page)
		pastPages[page] = i
	}

	return true
}
