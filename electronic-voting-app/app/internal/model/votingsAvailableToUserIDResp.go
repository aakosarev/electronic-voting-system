package model

type VotingsAvailableToUserIDResp struct {
	AvailableVotings []*AvailableVoting `json:"votings"`
}
