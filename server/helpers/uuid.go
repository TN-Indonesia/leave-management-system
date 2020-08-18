package helpers

import (
	uuid "github.com/satori/go.uuid"
)

// GetUUID ...
func GetUUID() string {
	uuidObject := uuid.NewV4()
	uid := uuidObject.String()
	return uid
}
