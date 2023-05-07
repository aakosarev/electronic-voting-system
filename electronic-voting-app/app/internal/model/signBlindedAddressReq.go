package model

type SignBlindedAddressReq struct {
	BlindedAddress []byte `json:"blinded_address"`
	VotingID       int32  `json:"voting_id"`
}
