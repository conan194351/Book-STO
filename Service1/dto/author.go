package dto

type FindBookByIdAuthorRequest struct {
	IdAuthor int `json:"idAuthor,omitempty"`
}

type Book struct {
	IdBook int64  `json:"IdBook,omitempty" gorm:"type:int"`
	Name   string `json:"Name,omitempty" gorm:"type:varchar(45)[]"`
}

type Books struct {
	Books []*Book `json:"books,omitempty"`
}
