package main

import (
	"file-manager/internal/window"
	"fmt"
	"os"
)

func main() {
	err := window.CreateWindow()
	if err != nil {
		fmt.Println("error: %s", err.Error())
		os.Exit(1)
	}
}
