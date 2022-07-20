package model

type Author struct {
	IdAuthor   int    `json:"idAuthor,omitempty"`
	Name       string `json:"name,omitempty"`
	NativeLand string `json:"NativeLand,omitempty"`
}
