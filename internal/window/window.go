package window

import (
	"file-manager/internal/colors"
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

	fmt.Print(colors.LightBlue, "=============YSNR's file manager=============", colors.Reset, "\n")
	fmt.Print("cwd: ", colors.Underline, cwd, "\n", colors.Reset)

	err = ShowDirs(cwd)
	if err != nil {
		return err
	}

	return nil
}

func ShowDirs(cwd string) error {
	dir, err := os.Open(cwd)
	if err != nil {
		return err
	}

	//пока и файлы и директории, сделать регулируемо дело пяти минут
	entities, err := dir.ReadDir(0) // 0 чтобы получить все сущности
	if err != nil {
		return err
	}

	fmt.Print("\n")
	for _, entity := range entities {
		info := entity.Name()
		if !entity.IsDir() {
			fileInfo, err := entity.Info()
			if err != nil {
				return err
			}

			//считаю размер
			size := float32(fileInfo.Size())
			ctr := 0
			for size > 1024 {
				size /= 1024
				ctr++
			}

			info += fmt.Sprintf(" <-%.1f%s", size, fileSizes[ctr])
		}
		fmt.Printf("\t%s\n", info)
	}

	return nil
}
