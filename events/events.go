package events

import (
	"errors"
	"time"

	"github.com/Kaktysfo/app/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID      string
	Title   string
	StartAt time.Time
}

func getNextID() string {
	return uuid.New().String()
}

func NewEvent(title string, dateStr string) (Event, error) {
	isValid := validation.IsValidateTitle(title)
	if isValid {
		dateParser, err := dateparse.ParseAny(dateStr)
		if err != nil {
			return Event{}, errors.New("неверный формат даты")
		}
		return Event{
			ID:      getNextID(),
			Title:   title,
			StartAt: dateParser,
		}, nil
	}
	return Event{}, errors.New("неправильный формат имени")
}
