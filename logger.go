package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("./vision/temp.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.Println("hello world")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(file)
	log.Println("hello world")
}
