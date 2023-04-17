package model

import "time"

type Voting struct {
	ID        int
	Title     string
	EndTime   int
	Address   string
	CreatedOn time.Time
}
