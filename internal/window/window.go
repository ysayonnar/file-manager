package window

import (
	"file-manager/internal/colors"
	commandparser "file-manager/internal/command-parser"
	"file-manager/internal/utils"
	"fmt"
	"os"
)

var (
	fileSizes = map[int]string{
		0: "b",
		1: "KB",
		2: "MB",
		3: "GB",
		4: "TB",
	}
)

func CreateWindow() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = RenderWindow(cwd)
	if err != nil {
		return err
	}
	return nil
}

func RenderWindow(cwd string) error {
	var commandError error
	for {
		fmt.Print(colors.LightBlue, "=============YSNR's file manager=============", colors.Reset, "\n")
		fmt.Print("cwd: ", colors.Underline, cwd, "\n", colors.Reset)

		catalog, err := ShowDirs(cwd)
		if err != nil {
			return err
		}

		if commandError != nil {
			fmt.Print("\n", colors.Red, "Error: ", commandError.Error(), colors.Reset, "\n")
		}

		var command string
		fmt.Print(colors.LightBlue, "\nCommand prompt: ", colors.Reset)
		fmt.Scanln(&command)
		cwd, commandError = commandparser.ParseCommand(cwd, command, catalog)

		utils.ClearScreen()
		utils.ClearStdin()
	}

	return nil
}

func ShowDirs(cwd string) (*commandparser.Catalog, error) {
	dir, err := os.Open(cwd)
	if err != nil {
		return nil, err
	}

	//пока и файлы и директории, сделать регулируемо дело пяти минут
	entities, err := dir.ReadDir(0) // 0 чтобы получить все сущности
	if err != nil {
		return nil, err
	}

	dirs := []os.DirEntry{}
	files := []os.DirEntry{}

	for _, entity := range entities {
		if entity.IsDir() {
			dirs = append(dirs, entity)
		} else {
			files = append(files, entity)
		}
	}

	fmt.Print("\n")
	for i, dir := range dirs {
		fmt.Print("- ", colors.Purple, fmt.Sprintf("%3d", i+1), " ", colors.Reset, dir.Name(), "\n")
	}

	for i, file := range files {
		info := file.Name()

		fileInfo, err := file.Info()
		if err != nil {
			return nil, err
		}

		//считаю размер
		size := float32(fileInfo.Size())
		ctr := 0
		for size > 1024 {
			size /= 1024
			ctr++
		}

		info += fmt.Sprintf(" <-%.1f%s", size, fileSizes[ctr])

		fmt.Print("- ", colors.Green, fmt.Sprintf("%3d", i+1), " ", colors.Reset, info, "\n")
	}

	return &commandparser.Catalog{Files: &files, Dirs: &dirs}, nil
}
