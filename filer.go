package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./temp.txt"

	file, err := os.Open(path)

	b := make([]byte, 1024)

	for {
		data, err := file.Read(b)
		if err != nil {
			fmt.Println("read file error")
			break
		}
		fmt.Println(string(b[:data]))
	}
}
