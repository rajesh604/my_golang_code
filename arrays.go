package main

import "fmt"

func main() {
	var grades [5]string
	for i := 0; i < 5; i++ {
		fmt.Scanf("%s", &grades[i])
	}
	fmt.Println(grades)
}
