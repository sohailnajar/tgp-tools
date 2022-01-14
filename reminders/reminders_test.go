package reminders_test

import (
	"bytes"
	"io"
	"reminders"
	"testing"
)

/*
reminder buy milk - write to text
reminder - print text from file
*/
func TestReminders(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("milk")
	fakeTerminal := &bytes.Buffer{}
	r, _ := reminders.NewReminder(
		reminders.WithInput(inputBuf),
		reminders.WithOutput(io.Writer(fakeTerminal)),
	)
	want := inputBuf.String()
	r.PrintReminders()
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}

// func TestNewReminders(t *testing.T) {
// 	t.Parallel()
// 	fakeTerminal := &bytes.Buffer{}
// 	f, _ := ioutil.ReadFile("testdata/test.txt")

// 	r := reminders.Reminder{
// 		Input:  bytes.NewReader(f),
// 		Output: fakeTerminal,
// 	}
// 	r.PrintReminders()
// 	want := string(f)
// 	got := fakeTerminal.String()
// 	if want != got {
// 		t.Errorf("want %q, got %q", want, got)
// 	}
// }

// func TestCreateReminders(t *testing.T) {
// 	t.Parallel()
// 	inputBuf := bytes.NewBufferString("hell")
// 	r := reminders.Reminder{
// 		Input: inputBuf,
// 	}
// 	r.NewReminder()

// }
