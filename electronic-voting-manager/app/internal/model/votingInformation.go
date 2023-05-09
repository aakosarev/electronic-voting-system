package model

import "time"

type VotingInformation struct {
	Title                  string
	NumberRegisteredVoters int64
	EndTime                time.Time
	Options                map[int64]*Option
}

type Option struct {
	Name        string
	NumberVotes int64
}
