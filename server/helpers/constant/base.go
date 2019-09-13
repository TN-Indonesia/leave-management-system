package constant

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//GetPass ...
func GetPass() string {
	GOPWD := os.Getenv("GOPWD")
	return GOPWD
}
