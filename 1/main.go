package main

import (
	"fmt"
	"math"
	"sort"
)

// Just brute force search this sucker and move on. I have better things to do.
func countContains(needle int, haystack []int) int {
	cnt := 0

	for i := range haystack {
		if haystack[i] == needle {
			cnt++
		}
	}

	return cnt
}

func main() {
	listA := []int{3, 4, 2, 1, 3, 3}
	listB := []int{4, 3, 5, 3, 9, 3}

	sort.Slice(listA, func(i, j int) bool {
		return listA[i] < listA[j]
	})

	sort.Slice(listB, func(i, j int) bool {
		return listB[i] < listB[j]
	})

	total := 0
	similarity := 0

	for i := range listA {
		a := listA[i]
		diff := int(math.Abs(float64(a - listB[i])))
		total += diff
		fmt.Printf("%d - %d = %d\n", a, listB[i], diff)

		similarity += a * countContains(a, listB)
	}

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Similarity: %d\n", similarity)
}
