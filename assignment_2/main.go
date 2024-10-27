package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file := os.Args[1]
	f, err := os.Open(file)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)

}
