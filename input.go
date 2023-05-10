package main

import "fmt"

func main() {
	var name string
	var is_muggle bool
	fmt.Println("What is your name?")
	fmt.Scan(&name, &is_muggle)
	fmt.Println("Hello, ", name, is_muggle)
}
