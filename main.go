package main

import (
	"fmt"
	"time"

	"github.com/Kaktysfo/app/calendar"
	"github.com/Kaktysfo/app/events"
)

func main() {
	e := events.Event{
		Title:   "Встреча",
		StartAt: time.Now(),
	}
	calendar.AddEvent("event1", e)
	fmt.Println("Календарь обновлён")
}
