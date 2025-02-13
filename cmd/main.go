package main

import (
	"file-manager/internal/window"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "help" {
			fmt.Println(`
	exit - Quit file manager
	od<index> - Open Directory
	of<index> - Open File
	back - move back by directory, as cd ..
	code - launch Visual Studio Code in current workind directory
	mkdir - create directory in cwd
	mkfile - create file in cwd
	dd<index> - delete dir
	df<index> - delete file
	rnd<index> - rename dir
	rnf<index> - rename file
			`)
		} else {
			fmt.Println("invalid prompt")
		}
		return
	}
	err := window.CreateWindow()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
}
