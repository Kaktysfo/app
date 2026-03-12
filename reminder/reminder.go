package reminder

import (
	"fmt"
	"time"
)

type Reminder struct {
	Message  string
	RemindAt time.Time
	Sent     bool
}

func NewReminder(message string, at time.Time) *Reminder {
	return &Reminder{
		Message:  message,
		RemindAt: at,
		Sent:     false,
	}
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	fmt.Println("Напоминание!", r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() { // нужно было самому написать?
}

// не понятно что значит может управлять свом состоянием
