package model

type Book struct {
	IdBook   int    `json:"_idBook,omitempty"`
	Ten_sach string `json:"Ten_sach,omitempty"`
	Ten_tg   string `json:"ten_tg,omitempty"`
	Cate     string `json:"cate,omitempty"`
}
