package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	fileContents, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, fileContents)

	// bs := make([]byte, 100)
	// fileContents.Read(bs)
	// fmt.Println(string(bs))

}
