package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./vision/test.txt", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = file.WriteString("Hello World")

	if err != nil {
		fmt.Println(err)
	}
}
