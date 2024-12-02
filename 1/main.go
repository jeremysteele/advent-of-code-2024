package main

import (
	"fmt"
	"math"
	"sort"
)

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

	for i := range listA {
		diff := int(math.Abs(float64(listA[i] - listB[i])))
		total += diff
		fmt.Printf("%d - %d = %d\n", listA[i], listB[i], diff)
	}

	fmt.Printf("Total: %d\n", total)
}
