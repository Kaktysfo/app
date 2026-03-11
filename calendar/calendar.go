package calendar

import (
	"errors"
	"fmt"

	"github.com/Kaktysfo/app/events"
)

//var calendarEvents = make(map[string]*events.Event)

var (
	EventError    = errors.New("cобытие с таким именем уже существует")
	addEventError = errors.New("ошибка добавления события")
	deleteError   = errors.New("ошибка удаления события")
	showError     = errors.New("список событий пуст")
)

type Calendar struct {
	calendarEvents map[string]*events.Event
}

func NewCalendar() *Calendar {
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
	}
}

func (c *Calendar) AddEvent(name, date string) (*events.Event, error) {
	event, err := events.NewEvent(name, date)
	if err != nil {
		return nil, addEventError
	}
	c.calendarEvents[event.ID] = event
	fmt.Println("Событие добавлено:", event.Title)
	return event, nil
}

func (c *Calendar) ShowEvents() error {
	if len(c.calendarEvents) == 0 {
		return showError
	}
	for _, event := range c.calendarEvents {
		fmt.Printf(
			"\nНазвание события: %s || Дата и время события: %s ||",
			event.Title,
			event.StartAt,
		)
	}

	fmt.Println("\n▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲▲")

	return nil
}

func (c *Calendar) isEventExist(id string) error {
	if _, ok := c.calendarEvents[id]; !ok {
		return EventError
	}
	return nil
}

func (c *Calendar) DeleteEvent(name string) error {
	err := c.isEventExist(name)
	if err != nil {
		return deleteError
	}
	delete(c.calendarEvents, name)
	fmt.Println("Успешно удалено!", name)
	return nil
}

func (c *Calendar) EditEvent(id, newTitle, dateStr string) error {
	e, exists := c.calendarEvents[id]
	if !exists {
		return errors.New("не удалось найти событие")
	}
	err := e.Update(newTitle, dateStr)
	return err
}

//func fullValidation(name, title string) error {
//	if _, ok := calendarEvents[name]; !ok {
//		return TitleError
//	}
//	if ok := validation.IsValidateTitle(title); !ok {
//		return errors.New("введен некорректно заголовок")
//	}
//	if calendarEvents[name].Title == title {
//		return errors.New("такой заголовок уже существует")
//	}
//	return nil
//}
