package main

import (
	"fmt"

	"github.com/Kaktysfo/app/calendar"
	"github.com/Kaktysfo/app/storage"
)

func main() {
	s := storage.NewStorage("calendar.json")
	c := calendar.NewCalendar(s)

	err := c.Load()
	if err != nil {
		if err.Error() == "такого файла не существует" {
			fmt.Println("Файл с событиями не найден. Будет создан новый файл при сохранении.")
		} else {
			fmt.Println("Ошибка загрузки данных:", err)
			return
		}
	}
	event1, err1 := c.AddEvent("jopa", "2025/06/12 14:00", "Высокий")
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(event1.Title, "added")
	}
	event2, err2 := c.AddEvent("One more meeting", "2025/06/12", "Высокий")
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(event2.Title, "added")
	}
	err32 := c.EditEvent(event2.ID, "Call", "2025/06/12 16:50", "Высокий")
	if err32 != nil {
		fmt.Println("Error:", err32)
	} else {
		fmt.Println("Event updated")
	}
	er12321 := c.DeleteEvent("Call") // не работает удаление
	if er12321 != nil {
		fmt.Println("Error", er12321)
	}
	errShow2 := c.ShowEvents()
	if errShow2 != nil {
		fmt.Println(errShow2)
	}
	defer c.Save() // не понял для чего defer
	fmt.Scanln()
}
