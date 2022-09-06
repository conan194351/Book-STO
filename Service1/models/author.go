package models

type Author struct {
	IdAuthor   int    `json:"idAuthor,omitempty"`
	Name       string `json:"name,omitempty"`
	NativeLand string `json:"NativeLand,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
}
