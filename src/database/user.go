package database

func FetchUsers() ([]User, error) {

	result, err := Database.Query("SELECT user.id, user.name, user.mail, user.birthday, user.wants_others_notified, user.wants_notifications FROM user;\n")
	defer close(result)
	if err != nil {
		return nil, err
	}

	var users []User

	for result.Next() {
		var user User
		err := result.Scan(&user.Id, &user.Name, &user.Mail, &user.Birthday, &user.WantsOthersNotified, &user.WantsNotifications)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
