package model

type RegistrationStatusesResp struct {
	Statuses map[int32]bool `json:"statuses"`
}
