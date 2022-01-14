package reminders

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type reminder struct {
	input  io.Reader
	output io.Writer
}

type option func(*reminder) error

func WithInput(input io.Reader) option {
	return func(r *reminder) error {
		if input == nil {
			return errors.New("nil input")
		}
		r.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(r *reminder) error {
		if output == nil {
			return errors.New("nil output")
		}
		r.output = output
		return nil
	}
}

func NewReminder(opts ...option) (reminder, error) {
	r := reminder{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(&r)
		if err != nil {
			return reminder{}, err
		}
	}
	return r, nil
}

func PrintReminders() {
	content, err := ioutil.ReadFile("reminders.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

}

func (r reminder) SaveReminders(args []string) {
	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	writer := bufio.NewWriter(file)
	if err != nil {
		log.Fatal(err)
	}

}

// func NewReminder() (Reminder, error) {
// 	f, err := os.Open("reminders.txt")
// 	if err != nil {
// 		return Reminder{}, err
// 	}
// 	defer f.Close()
// 	w := bufio.NewReader(f)
// 	return Reminder{
// 		Input:  w,
// 		Output: os.Stdout,
// 	}, nil
// }

// func (r Reminder) PrintReminders() {
// 	io.Copy(r.Output, r.Input)
// }

// func (r Reminder) AddReminder(args []string) {
// 	f, _ := os.OpenFile("reminders.txt", os.O_APPEND|os.O_WRONLY, 0644)
// 	defer f.Close()
// 	for _, item := range args {
// 		f.Write([]byte(item))
// 	}
// }
