package main

import (
	"file-manager/internal/window"
	"fmt"
	"os"
)

func main() {
	err := window.CreateWindow()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}
