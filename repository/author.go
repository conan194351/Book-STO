package repository

import (
	"book-sto/errs"
	"book-sto/model"
	"database/sql"
)

type AuthorRepository interface {
	List() ([]model.Author, *errs.AppError)
	Create(model.Author) (model.Author, *errs.AppError)
	SearchAuthor(req string) ([]model.Author, *errs.AppError)
}

type DefaultAuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {

	return DefaultAuthorRepository{

		db: db,
	}
}

func (a DefaultAuthorRepository) List() ([]model.Author, *errs.AppError) {

	res, err := a.db.Query("SELECT * FROM author")

	if err != nil {

		return nil, errs.ErrorGetData()
	}

	var authors []model.Author
	var author model.Author
	for res.Next() {
		err = res.Scan(&author.IdAuthor, &author.Name, &author.NativeLand)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a DefaultAuthorRepository) Create(author model.Author) (model.Author, *errs.AppError) {

	newAuthor, err := a.db.Prepare("INSERT INTO longphu.author(Name, NativeLand) VALUES (? ,?) ")
	if err != nil {

		return author, errs.ErrorInsertData()
	}
	newAuthor.Exec(author.Name, author.NativeLand)
	return author, nil
}

func (a DefaultAuthorRepository) SearchAuthor(req string) ([]model.Author, *errs.AppError) {
	var authors []model.Author
	if req == "" {
		return authors, nil
	}
	bodyString := "%" + req + "%"
	res, err := a.db.Query("SELECT author.Name, author.NativeLand FROM longphu.author as author WHERE author.Name LIKE ?", bodyString)

	if err != nil {

		return nil, errs.ErrorGetData()
	}
	for res.Next() {

		var author model.Author
		err = res.Scan(&author.Name, &author.NativeLand)
		if err != nil {

			return nil, errs.ErrorReadData()
		}

		authors = append(authors, author)
	}

	return authors, nil
}
