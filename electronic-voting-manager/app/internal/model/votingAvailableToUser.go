package model

import "time"

type VotingAvailableToUser struct {
	UserID    int
	VotingID  int
	CreatedOn time.Time
	Title     string
	Address   string
}
