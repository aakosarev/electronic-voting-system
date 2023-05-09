package model

type DemoCastAVoteReq struct {
	VotingID int32  `json:"voting_id"`
	Address  string `json:"address"`
	Index    int32  `json:"index"`
}
