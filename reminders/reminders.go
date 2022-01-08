package reminders

import (
	"bufio"
	"io"
	"os"
)

type Reminder struct {
	Input  io.Reader
	Output io.Writer
}

func NewReminder() (Reminder, error) {
	f, err := os.Open("reminders.txt")
	if err != nil {
		return Reminder{}, err
	}
	defer f.Close()
	w := bufio.NewReader(f)
	return Reminder{
		Input:  w,
		Output: os.Stdout,
	}, nil
}

func (r Reminder) PrintReminders() {
	io.Copy(r.Output, r.Input)
}

func (r Reminder) AddReminder(args []string) {
	f, _ := os.OpenFile("reminders.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	for _, item := range args {
		f.Write([]byte(item))
	}
}
