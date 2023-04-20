package model

import (
	"time"
)

type GetAllVotingsResp struct {
	VotingID    int32     `json:"voting_id"`
	VotingTitle string    `json:"voting_title"`
	EndTime     int64     `json:"end_time"`
	Address     string    `json:"address"`
	CreatedOn   time.Time `json:"created_on"`
}
