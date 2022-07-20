package dto

import "book-sto/model"

type GetAllAuthorResponse struct {
	Authors []model.Author
}

type CreateAuthorResponse struct {
	Author model.Author
}

type CreateAutherRequest struct {
	Name       string `json:"Name"`
	NativeLand string `json:"NativeLand"`
}

type SearchAuthorRequest struct {
	Name string `json:"Name"`
}

type SearchAuthorResponse struct {
	Authors []model.Author
}
