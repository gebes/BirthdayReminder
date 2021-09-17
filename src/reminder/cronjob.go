package reminder

import (
	"gebes.io/go-birthdayreminder/src/database"
	"gebes.io/go-birthdayreminder/src/mail"
	"github.com/robfig/cron"
	"log"
	"strconv"
	"strings"
	"time"
)

func StartCronjob() error {

	log.Println("Initializing CronJob")
	location, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	c := cron.NewWithLocation(location)
	err = c.AddFunc("0 0 6 * * *", task)
	if err != nil {
		return err
	}
	c.Start()

	select {}

}

func task() {
	err := remindEveryone()
	if err != nil {
		log.Println(err)
	}
}

func remindEveryone() error {

	log.Println("Fetching users")
	users, err := database.FetchUsers()

	if err != nil {
		return err
	}

	for _, user := range users {

		if !user.WantsNotifications {
			continue
		}

		var birthdayPeople []database.User
		for _, userToCheck := range users {
			if userToCheck.Id == user.Id || !user.WantsOthersNotified {
				continue
			}

			if hasBirthdayToday(&userToCheck) {
				birthdayPeople = append(birthdayPeople, userToCheck)
			}

		}

		if len(birthdayPeople) == 0 {
			continue
		}

		ageMap := map[int][]database.User{}
		for _, birthdayUser := range birthdayPeople {
			age := getAge(&birthdayUser)
			list, _ := ageMap[age]
			list = append(list, birthdayUser)
			ageMap[age] = list
		}

		birthdays := ""
		for key, value := range ageMap {
			birthdays += ageLineFromList(value, key)
		}

		var replace = map[string]string{}
		replace["Name"] = user.Name
		replace["Birthdays"] = birthdays

		content, err := mail.ParseTemplate("./mails/email.html", replace)
		if err != nil {
			return err
		}

		_, err = mail.SendGridMail(&mail.Mail{
			Receiver: user.Mail,
			Subject:  subjectFromList(birthdayPeople),
			Body:     *content,
		})
		if err != nil {
			return err
		}
		log.Println("Sent mail to " + user.Mail)

	}

	return nil
}

func hasBirthdayToday(user *database.User) bool {
	today := time.Now()
	todayDay := today.Day()
	todayMonth := int(today.Month())

	day, month, _ := parseUserBirthday(user)

	return day == todayDay && month == todayMonth
}

func getAge(user *database.User) int {
	currentYear := time.Now().Year()
	_, _, year := parseUserBirthday(user)
	return currentYear - year
}

func subjectFromList(users []database.User) string {
	if len(users) > 1 {
		return namesListed(users) + " haben heute Geburtstag"
	} else {
		return namesListed(users) + " hat heute Geburtstag"
	}
}

func ageLineFromList(users []database.User, age int) string {
	if len(users) > 1 {
		return namesListed(users) + " werden heute " + strconv.Itoa(age) + " Jahre alt. "
	} else {
		return namesListed(users) + " wird heute " + strconv.Itoa(age) + " Jahre alt. "
	}
}

func parseUserBirthday(user *database.User) (int, int, int) {
	split := strings.Split(user.Birthday, ".")
	day, _ := strconv.Atoi(split[0])
	month, _ := strconv.Atoi(split[1])
	year, _ := strconv.Atoi(split[2])
	return day, month, year
}

func namesListed(users []database.User) string {
	if len(users) == 1 {
		return users[0].Name
	} else if len(users) > 1 {
		result := ""

		for i, user := range users {
			result += user.Name
			if i == len(users)-2 {
				result += " und "
			} else if i != len(users)-1 {
				result += ", "
			}
		}

		return result
	} else {
		return "Niemand"
	}
}
