package helpers

import "net/http"

//CheckValidImage ...
func CheckValidImage(buffer []byte) bool {
	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" && contentType != "image/png" {
		return false
	}

	return true
}
