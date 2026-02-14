package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	JWT_SECRET = ""
	PORT       = ""
)

func Load() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ENV NOT LOADED")
		return err
	}
	JWT_SECRET = os.Getenv("JWT_SECRET")
	PORT = os.Getenv("PORT")
	fmt.Println("ENV LOADED")
	return nil
}
