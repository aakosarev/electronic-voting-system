package model

import "time"

type AvailableVoting struct {
	UserID    int32     `json:"user_id"`
	VotingID  int32     `json:"voting_id"`
	CreatedOn time.Time `json:"created_on"`
	Title     string    `json:"title"`
	Address   string    `json:"address"`
}
