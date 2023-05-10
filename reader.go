package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println(err)
		// log.Fatal(err)
	}
}
