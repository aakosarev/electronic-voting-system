package model

import "time"

type Voting struct {
	ID        int32
	Title     string
	EndTime   int64
	Address   string
	CreatedOn time.Time
}
