package main

import (
	"fmt"
)

func main() {
	var a string
	var b int

	fmt.Println("Enter a string and an integer:")

	count, err := fmt.Scan(&a, &b)

	fmt.Println(a, b)
	fmt.Println("Number of inputs:", count)
	fmt.Println("Error:", err)
}
