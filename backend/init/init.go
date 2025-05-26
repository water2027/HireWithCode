package init

import (
	"backend/config"
	"backend/database"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}
	config.InitConfig()
	database.InitDatabase()
} 