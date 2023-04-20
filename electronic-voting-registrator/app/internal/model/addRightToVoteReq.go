package model

type AddRightToVoteReq struct {
	UserID   int32 `json:"user_id"`
	VotingID int32 `json:"voting_id"`
}
