package utils

func IsNumber(b byte) bool {
	s := string(b)
	if s == "0" || s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" || s == "8" || s == "9" {
		return true
	}
	return false
}
