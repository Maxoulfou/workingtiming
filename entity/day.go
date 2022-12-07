package entity

import "time"

type day struct {
	ID         int16
	Name       string
	StartTime  time.Time
	EndTime    time.Time
	StartBreak time.Time
	StopBreak  time.Time
}
