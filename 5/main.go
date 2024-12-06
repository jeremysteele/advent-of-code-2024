package main

import (
	"os"
	"strings"
)

func main() {
	ruleData, err := os.ReadFile("5/rules.txt")

	if err != nil {
		panic(err)
	}

	updateData, err := os.ReadFile("5/updates.txt")

	if err != nil {
		panic(err)
	}

	rules := make(map[string][]string)

	for _, line := range strings.Split(string(ruleData), "\n") {
		rule := strings.Split(line, "|")

		if _, ok := rules[rule[0]]; !ok {
			rules[rule[0]] = []string{}
		}

		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	var updates [][]string
	for _, line := range strings.Split(string(updateData), "\n") {
		updates = append(updates, strings.Split(line, ","))
	}

	part2Process(rules, updates)
}
