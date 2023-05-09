package model

import "time"

type VotingInformationResp struct {
	Title                  string            `json:"title"`
	NumberRegisteredVoters int64             `json:"number_registered_voters"`
	EndTime                time.Time         `json:"end_time"`
	Options                map[int64]*Option `json:"options"`
}

type Option struct {
	Name        string `json:"name"`
	NumberVotes int64  `json:"number_votes"`
}
