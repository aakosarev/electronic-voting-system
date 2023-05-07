package model

type LoginReq struct {
	UserID   int32  `json:"user_id"`
	Password string `json:"password"`
}
