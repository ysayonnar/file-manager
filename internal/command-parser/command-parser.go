package commandparser

import (
	"file-manager/internal/colors"
	"file-manager/internal/utils"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	MAX_FILE_SIZE_TO_READ = 5120
	PATH_TO_VSCODE        = `C:\Users\ysayo\AppData\Local\Programs\Microsoft VS Code\Code.exe`
	commands              = map[string]func(commandInput CommandInput) (wd string, err error){
		"od":     OpenDir,    // opens directory
		"of":     OpenFile,   //opens file
		"exit":   Exit,       // exit file-manager
		"back":   BackDir,    // back on tree, as `cd ..`
		"code":   LaunchCode, // launch vs code in the cwd
		"mkdir":  MakeDir,
		"mkfile": MakeFile,
		"dd":     DeleteDir,
		"df":     DeleteFile,
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
	utils.ClearScreen()
	os.Exit(1)
	return "", nil
}

func BackDir(commandInput CommandInput) (wd string, err error) {
	wd = filepath.Join(commandInput.cwd, "..")
	return wd, nil
}

func OpenFile(commandInput CommandInput) (wd string, err error) {
	index := commandInput.index
	if index > len(*commandInput.catalog.Files) || index < 1 {
		return commandInput.cwd, fmt.Errorf("invalid index of file")
	}
	file := (*commandInput.catalog.Files)[index-1]

	info, err := file.Info()
	if err != nil {
		return commandInput.cwd, fmt.Errorf("impossible to read")
	}
	if info.Size() > int64(MAX_FILE_SIZE_TO_READ) {
		return commandInput.cwd, fmt.Errorf("max file size reached -> %d bytes", MAX_FILE_SIZE_TO_READ)
	}
	filePath := filepath.Join(commandInput.cwd, file.Name())
	data, err := os.ReadFile(filePath)
	if err != nil {
		return commandInput.cwd, fmt.Errorf("unexpected error while readind file")
	}

	lines := strings.Split(string(data), "\n")

	utils.ClearScreen()
	fmt.Print(colors.Green, file.Name(), colors.Reset, "\n")
	fmt.Print(colors.LightBlue, "--------- READ ONLY ---------\n")
	for i, line := range lines {
		fmt.Print(colors.Purple, fmt.Sprintf("%3d  ", i+1), colors.Reset, line, "\n")
	}

	var exitSignal string
	fmt.Print("\n", colors.LightBlue, colors.Underline, "Use `Enter` to exit read mode", colors.Reset, "\n")
	fmt.Scanln(&exitSignal)

	return commandInput.cwd, nil
}

func LaunchCode(commandInput CommandInput) (wd string, err error) {
	cmd := exec.Command(PATH_TO_VSCODE, commandInput.cwd)
	err = cmd.Run()
	if err != nil {
		return commandInput.cwd, fmt.Errorf("path to code is invalid, edit config")
	}
	return commandInput.cwd, nil
}

func MakeDir(commandInput CommandInput) (wd string, err error) {
	fmt.Print(colors.Green, "Name of the directory: ", colors.Reset)
	var dirName string
	for dirName == "" {
		fmt.Scanln(&dirName)
	}

	fmt.Print("\n", colors.Green, "Should it be available to read?(y - yes):  ", colors.Reset)
	var readMode string
	fmt.Scanln(&readMode)

	fmt.Print("\n", colors.Green, "Should it be available to write?(y - yes):  ", colors.Reset)
	var writeMode string
	fmt.Scanln(&writeMode)

	permissions := 0
	if readMode == "y" {
		permissions += 4
	}
	if writeMode == "y" {
		permissions += 2
	}
	if permissions == 0 {
		return commandInput.cwd, fmt.Errorf("directory can't be created without any permissions")
	}

	newDirPath := filepath.Join(commandInput.cwd, dirName)
	err = os.Mkdir(newDirPath, os.FileMode(permissions))
	return commandInput.cwd, err
}

func MakeFile(commandInput CommandInput) (wd string, err error) {
	fmt.Print(colors.Green, "Name of the file with extension: ", colors.Reset)
	var fileName string
	for fileName == "" {
		fmt.Scanln(&fileName)
	}

	filePath := filepath.Join(commandInput.cwd, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return commandInput.cwd, err
	}
	defer file.Close()

	return commandInput.cwd, nil
}

func DeleteFile(commandInput CommandInput) (wd string, err error) {
	index := commandInput.index
	if index > len(*commandInput.catalog.Files) || index < 1 {
		return commandInput.cwd, fmt.Errorf("invalid index of file")
	}
	file := (*commandInput.catalog.Files)[index-1]
	filePath := filepath.Join(commandInput.cwd, file.Name())

	return commandInput.cwd, os.Remove(filePath)
}

func DeleteDir(commandInput CommandInput) (wd string, err error) {
	index := commandInput.index
	if index > len(*commandInput.catalog.Dirs) || index < 1 {
		return commandInput.cwd, fmt.Errorf("invalid index of dir")
	}
	dir := (*commandInput.catalog.Dirs)[index-1]
	dirPath := filepath.Join(commandInput.cwd, dir.Name())
	return commandInput.cwd, os.Remove(dirPath)
}
