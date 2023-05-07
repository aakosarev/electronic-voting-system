package model

type RegisterAddressReq struct {
	Address       string `json:"address"`
	SignedAddress []byte `json:"signed_address"`
	VotingID      int32  `json:"voting_id"`
}
