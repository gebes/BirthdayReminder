package main

import (
	"gebes.io/go-birthdayreminder/src/database"
	"gebes.io/go-birthdayreminder/src/env"
	"gebes.io/go-birthdayreminder/src/reminder"
)

func main() {

	if err := env.LoadEnvironment(); err != nil {
		panic(err)
	}

	database.InitDatabase()

	if err := reminder.StartCronjob(); err != nil {
		panic(err)
	}

}
