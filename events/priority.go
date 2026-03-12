package events

import "errors"

type Priority string

const (
	PriorityHigh   Priority = "Высокий"
	PriorityMedium Priority = "Средний"
	PriorityLow    Priority = "Низкий"
)

func (p Priority) Validate() error {
	switch p {
	case PriorityHigh, PriorityMedium, PriorityLow:
		return nil
	default:
		return errors.New("неверный приоритет")
	}
}
