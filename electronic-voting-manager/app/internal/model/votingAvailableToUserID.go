package model

import "time"

type VotingAvailableToUserID struct {
	UserID    int32
	VotingID  int32
	CreatedOn time.Time
	Title     string
	Address   string
}
