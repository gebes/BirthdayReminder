package database

type User struct {
	Id                  int
	Name                string
	Mail                string
	Birthday            string
	WantsOthersNotified bool
	WantsNotifications  bool
}
