package entity

import "time"

type Holiday struct {
	ID          int16
	Description string
	Start       time.Time
	End         time.Time
}
