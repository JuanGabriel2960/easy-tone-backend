package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error reading environment variables.")
		os.Exit(1)
	}
}
