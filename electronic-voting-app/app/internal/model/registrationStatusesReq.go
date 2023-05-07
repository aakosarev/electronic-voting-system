package model

type RegistrationStatusesReq struct {
	Addresses map[int32]string `json:"addresses"`
}
