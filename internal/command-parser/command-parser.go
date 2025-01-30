package commandparser

import (
	"file-manager/internal/utils"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

var (
	commands = map[string]func(commandInput CommandInput) (wd string, err error){
		"od":   OpenDir,
		"exit": Exit,
		"back": BackDir,
	}
)

type CommandInput struct {
	cwd     string
	catalog *Catalog
	index   int
}

type Catalog struct {
	Files *[]fs.DirEntry
	Dirs  *[]fs.DirEntry
}

func ParseCommand(cwd string, command string, catalog *Catalog) (wd string, err error) {
	if command == "" {
		return cwd, nil
	}
	commandCode := []byte{}
	number := []byte{}
	for _, symbol := range command {
		if utils.IsNumber(byte(symbol)) {
			number = append(number, byte(symbol))
		} else {
			commandCode = append(commandCode, byte(symbol))
		}
	}

	if _, ok := commands[string(commandCode)]; !ok {
		return cwd, fmt.Errorf("unknown command `%s`", command)
	}

	f := commands[string(commandCode)]
	index, err := strconv.Atoi(string(number))
	if err != nil && len(number) > 0 {
		return cwd, fmt.Errorf("invalid index `%s`", command)
	}

	return f(CommandInput{cwd: cwd, catalog: catalog, index: index})
}

func OpenDir(commandInput CommandInput) (wd string, err error) {
	index := commandInput.index
	if index > len(*commandInput.catalog.Dirs) || index < 1 {
		return commandInput.cwd, fmt.Errorf("invalid index of dir")
	}
	dir := (*commandInput.catalog.Dirs)[index-1]
	wd = filepath.Join(commandInput.cwd, dir.Name())
	return wd, nil
}

func Exit(commandInput CommandInput) (wd string, err error) {
	fmt.Println("Bye-bye!")
	os.Exit(1)
	return "", nil
}

func BackDir(CommandInput CommandInput) (wd string, err error) {
	wd = filepath.Join(CommandInput.cwd, "..")
	return wd, nil
}
