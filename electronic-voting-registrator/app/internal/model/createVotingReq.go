package model

import "time"

type CreateVotingReq struct {
	Title   string    `json:"title"`
	Options string    `json:"options"`
	EndTime time.Time `json:"end_time"`
}
