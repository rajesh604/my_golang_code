package main

import (
	"fmt"
	"sort"
)

func main() {
	// vars := []int{5, 2, 6, 3, 1, 4}
	// sort.Ints(vars)
	// fmt.Println(vars)

	// vars := []string{"c", "a", "b"}
	// sort.Strings(vars)
	// fmt.Println(vars)

	vars := []string{"Learning", "Golang", "on", "Kodekloud"}
	sort.Strings(vars)
	fmt.Println(vars)
}
