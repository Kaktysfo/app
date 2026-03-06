package main

import (
	"fmt"

	"github.com/Kaktysfo/app/calendar"
	"github.com/Kaktysfo/app/events"
)

func main() {
	event1, err := events.NewEvent("День рождения", "2026/09/03 13:33")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
	} else {
		calendar.AddEvent("event 1", event1)
	}
	event2, err2 := events.NewEvent("Уйти в запой", "03/09/2026")
	if err2 != nil {
		fmt.Println("Ошибка создания события:", err2)
	} else {
		calendar.AddEvent("event 2", event2)
	}
	event3, err3 := events.NewEvent("popopooop", "03/09/2026 15:15")
	if err3 != nil {
		fmt.Println("Ошибка создания события:", err3)
	} else {
		calendar.AddEvent("event 3", event3)
	}
	err11 := calendar.DeleteEvent("event 3")
	if err11 != nil {
		fmt.Println(err11)
	}
	err12 := calendar.EditEvent("event 2", "jopa123321", "12/12/2026 11:15")
	if err12 != nil {
		fmt.Println(err12)
	}
	calendar.ShowEvents()
	_, errScan := fmt.Scanln()
	if errScan != nil {
		return
	}
}
