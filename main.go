package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	fileContents, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 100)
	fileContents.Read(bs)
	fmt.Println(string(bs))

}
