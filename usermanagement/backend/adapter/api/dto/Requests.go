package dto

import model "github.com/D4ykoo/travelplatform-case-m2/usermanagement/domain/model"

type CreateUserRequest struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserRequest struct {
	Username    string `json:"username"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}

type ResetPasswordRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ue CreateUserRequest) ToUser() model.User {
	return model.User{
		Username:  ue.Username,
		Firstname: ue.Firstname,
		Lastname:  ue.Lastname,
		Email:     ue.Email,
		Password:  ue.Password,
	}
}
