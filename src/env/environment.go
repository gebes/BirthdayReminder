package env

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var (
	MySqlDatabase string

	MailServerAddress  string
	MailServerPort     int
	MailServerUsername string
	MailServerPassword string

	SendgridApiKey string
)

func LoadEnvironment() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}
	mailServerPort, err := strconv.Atoi(os.Getenv("MAIL_SERVER_PORT"))
	if err != nil {
		return err
	}
	MySqlDatabase = os.Getenv("MYSQL_DATABASE")
	MailServerAddress = os.Getenv("MAIL_SERVER_ADDRESS")
	MailServerPort = mailServerPort
	MailServerUsername = os.Getenv("MAIL_USERNAME")
	MailServerPassword = os.Getenv("MAIL_PASSWORD")
	SendgridApiKey = os.Getenv("SENDGRID_API_KEY")

	return nil
}


