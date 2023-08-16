package dto

import "Leilei-platform/model"

type ProjectResponse struct {
	Total int64           `json:"total"`
	List  []model.Project `json:"list"`
}

type ProjectUsers struct {
	Total int64           `json:"total"`
	List  []model.Project `json:"list"`
}
