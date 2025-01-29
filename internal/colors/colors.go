package colors

import "fmt"

var (
	Black     = "\033[30m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Purple    = "\033[35m"
	LightBlue = "\033[36m"
	White     = "\033[37m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Reset     = "\033[0m"
)

func TestAllColors(msg string) {
	fmt.Println(Black, msg)
	fmt.Println(Red, msg)
	fmt.Println(Green, msg)
	fmt.Println(Yellow, msg)
	fmt.Println(Black, msg)
	fmt.Println(Purple, msg)
	fmt.Println(LightBlue, msg)
	fmt.Println(White, msg)
	fmt.Println(Bold, msg)
	fmt.Println(Underline, msg)
}
