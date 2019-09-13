package constant

import (
	"os"
)

//GetPass ...
func GetPass() string {
	GOPWD := os.Getenv("GOPWD")
	return GOPWD
}

//GetPassReset ...
func GetPassReset() string {
	GOPWDRESET := os.Getenv("GOPWDRESET")
	return GOPWDRESET
}
