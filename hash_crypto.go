package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func encodeWithMD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func main() {
	var password string
	fmt.Scanln(&password)
	fmt.Println("Password encrypted to", encodeWithMD5(password))
}
