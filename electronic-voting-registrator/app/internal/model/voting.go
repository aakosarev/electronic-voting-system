package model

import "time"

type Voting struct {
	ID        int32     `json:"id"`
	Title     string    `json:"title"`
	EndTime   int64     `json:"end_time"`
	Address   string    `json:"address"`
	CreatedOn time.Time `json:"created_on"`
}
