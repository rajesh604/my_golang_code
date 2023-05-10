package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	buf := make([]byte, 4)

	for {
		n, err := r.Read(buf)
		fmt.Println(buf[:n], err)
		fmt.Println(string(buf[:n]))
		if err != nil {
			fmt.Println("breaking out")
			break
		}
	}
}
