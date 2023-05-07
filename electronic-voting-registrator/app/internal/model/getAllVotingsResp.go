package model

type GetAllVotingsResp struct {
	Votings []*Voting `json:"votings"`
}
