package dto

import "Leilei-platform/model"

type RoleDTO struct {
	Total int64        `json:"total"`
	List  []model.Role `json:"list"`
}

type RoleUser struct {
	Total int64        `json:"total"`
	List  []model.User `json:"list"`
}
