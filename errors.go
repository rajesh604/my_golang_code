package main

import (
	"errors"
	"fmt"
)

func process(i int) err {
	if i%2 == 0 {
		return errors.New("Only odd numbers are allowed")
	}
	return nil
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	err := process(3)
	checkError(err)
	err = process(2)
	checkError(err)
}
