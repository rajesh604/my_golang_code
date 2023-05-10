package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("./vision/temp.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.Println("hello world")
	if err != nil {
		fmt.Println(err)
		return
	}
	// log.SetOutput(file)
	// log.Println("hello world")
	logrus.SetOutput(file)
	logrus.Println("hello world")
}
