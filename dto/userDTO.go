package dto

import "Leilei-platform/model"

type UserDTO struct {
	model.User
	Permissions []string `json:"permissions"`
	Token       string   `json:"token"`
}

type SwitchProjectDTO struct {
	LastProject string `json:"lastProject"`
	ID          string `json:"id"`
}

type UpdatePasswordDTO struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
	UserId      string `json:"userId"`
}
type UpdateInfoDTO struct {
	Username string `json:"userName"`
	Email    string `json:"email"`
}
