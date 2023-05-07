package model

import "time"

type InformationAboutVoting struct {
	Title                  string
	OptionsCompleted       bool
	NumberRegisteredVoters int64
	EndTime                time.Time
	Options                map[int64]*Option
}

type Option struct {
	Name        string
	NumberVotes int64
}
