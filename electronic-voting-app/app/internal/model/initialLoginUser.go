package model

type InitialLoginUser struct {
	FirstName         string `json:"first_name"`
	SecondName        string `json:"second_name"`
	Email             string `json:"email"`
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	ReenteredPassword string `json:"reentered_password"`
}
