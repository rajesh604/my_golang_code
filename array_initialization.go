package main

import "fmt"

func main() {
	var fruits [2]string = [2]string{"apple", "banana"}
	fmt.Println(fruits)
	marks := [3]int{1, 2, 3}
	fmt.Println(marks)
	names := [...]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
}
