package calendar

import (
	"errors"
	"fmt"

	"github.com/Kaktysfo/app/events"
	"github.com/Kaktysfo/app/validation"
	"github.com/araddon/dateparse"
)

var eventsMap = make(map[string]events.Event)
var TitlteErr = errors.New("cобытие с таким именем уже существует")

func AddEvent(key string, e events.Event) {
	eventsMap[key] = e
	fmt.Println("Событие добавлено: ", e.Title)
}

func ShowEvents() {
	fmt.Println("\nВсе события в календаре: ")
	fmt.Println("▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼▼")
	for _, event := range eventsMap {
		fmt.Printf("\nНазвание события:  %s || Дата и время события:  %s ||", event.Title, event.StartAt)
	}
	fmt.Println("\n \n▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲")
}

func isTitleExist(title string) error {
	if _, ok := eventsMap[title]; !ok {
		return TitlteErr
	}
	return nil
}

func DeleteEvent(key string) error {
	err := isTitleExist(key)
	if err != nil {
		return err
	}
	delete(eventsMap, key)
	fmt.Println("Успешно удалено!", key)
	return nil
}

func EditEvent(name, newTitle, dateStr string) error {
	date, dateErr := dateparse.ParseAny(dateStr)
	if dateErr != nil {
		return dateErr
	}
	err := fullValidation(name, newTitle)
	if err != nil {
		return err
	}
	eventsMap[name] = events.Event{
		Title:   newTitle,
		StartAt: date,
	}
	fmt.Println("Успешно изменено!", name)
	return nil
}

func fullValidation(name, title string) error {
	if _, ok := eventsMap[name]; !ok {
		return TitlteErr
	}
	if ok := validation.IsValidateTitle(title); !ok {
		return errors.New("введен некорректно заголовок")
	}
	if eventsMap[name].Title == title {
		return errors.New("такой заголовок уже существует")
	}
	return nil
}
