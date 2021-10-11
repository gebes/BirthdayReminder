# BirthdayReminder
[![GoReportCard](https://goreportcard.com/badge/github.com/gebes/go-birthdayreminder)](https://goreportcard.com/report/github.com/gebes/go-birthdayreminder)

Mail birthday reminder for our school class.

![](https://i.imgur.com/GCtIrIs.png)

## Functionality
BirthdayReminder checks every day at 6:00 am if someone has a birthday. If so, then BirthdayReminder generates a mail and sends it to everyone except the birthday person. If multiple people have a birthday on the same day, then BirthdayReminder will adapt the message to be grammatically correct.

## Requirements

### Setup a MySQL Database
Create a `user` table.
```mysql
CREATE TABLE IF NOT EXISTS user(
    id SERIAL,
    name VARCHAR(64),
    mail VARCHAR(128),
    birthday VARCHAR(10),
    wants_others_notified BOOL,
    wants_notifications BOOL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Create some entries.
```mysql
INSERT INTO user(name, mail, birthday, wants_others_notified, wants_notifications) 
VALUES ("Christoph", "mail@mail.com", "01.01.1970", true, true), 

### Get a Sendgrid API Key
Create a Sendgrid Account and get your API Key

### .env

```dotenv
MYSQL_DATABASE=user:password@tcp(host:port)/birthdayreminder
SENDGRID_API_KEY=SG.key123abc
```

### Deploy
Use the Dockerfile to deploy the BirthdayReminder on your server.
