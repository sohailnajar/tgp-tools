package main

import (
	"os"
	"reminders"
)

func main() {
	r, _ := reminders.NewReminder()
	if len(os.Args) > 1 {
		r.SaveReminders(os.Args[1:])
	} else {
		reminders.PrintReminders()
	}

}
