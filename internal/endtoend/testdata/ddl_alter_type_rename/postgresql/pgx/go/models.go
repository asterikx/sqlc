// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package querytest

import (
	"fmt"
)

type NewEvent string

const (
	NewEventSTART NewEvent = "START"
	NewEventSTOP  NewEvent = "STOP"
)

func (e *NewEvent) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NewEvent(s)
	case string:
		*e = NewEvent(s)
	default:
		return fmt.Errorf("unsupported scan type for NewEvent: %T", src)
	}
	return nil
}

type LogLine struct {
	ID     int64
	Status NewEvent
}
