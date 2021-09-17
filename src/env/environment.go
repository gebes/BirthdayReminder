package env

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	MySqlDatabase string



	SendgridApiKey string
)

func LoadEnvironment() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	MySqlDatabase = os.Getenv("MYSQL_DATABASE")
	SendgridApiKey = os.Getenv("SENDGRID_API_KEY")

	return nil
}


