package model

type InitialLoginReq struct {
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Email             string `json:"email"`
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	ReenteredPassword string `json:"reentered_password"`
}
