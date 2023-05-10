package main

import "fmt"

func main() {
	for index, element := range [3]string{"a", "b", "c"} {
		fmt.Println(index, element)
	}
}
