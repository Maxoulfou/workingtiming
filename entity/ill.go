package entity

import "time"

type Ill struct {
	ID                     int16
	StartDate              time.Time
	EndDate                time.Time
	Reason                 string
	HaveIllnessCertificate bool
}
