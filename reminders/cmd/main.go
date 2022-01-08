package main

import (
	"os"
	"reminders"
)

func main() {
	r, _ := reminders.NewReminder()
	if os.Args == nil {
		r.PrintReminders()
	}
	r.AddReminder(os.Args[1:])

}
