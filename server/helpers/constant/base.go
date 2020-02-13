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

// GetClientURL ...
func GetClientURL() string {
	protocol := os.Getenv("client_protocol")
	addr := os.Getenv("client_addr")
	port := os.Getenv("client_port")

	return protocol + "://" + addr + ":" + port
}
