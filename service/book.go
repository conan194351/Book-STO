package service

import (
	"book-sto/config"
	"book-sto/model"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Book model.Book

var db = config.DbConn()

func IndexBook(response http.ResponseWriter, request *http.Request) {
	selDB, err := db.Query("SELECT book.idBook,book.Ten_sach FROM longphu.book as book")
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "500"))
		return
	}
	defer selDB.Close()
	book := Book{}
	res := []Book{}
	for selDB.Next() {
		var idBook int
		var ten_sach string
		err = selDB.Scan(&idBook, &ten_sach)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
			return
		}
		book.IdBook = idBook
		book.Ten_sach = ten_sach
		res = append(res, book)
	}
	json.NewEncoder(response).Encode(res)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	var book Book
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "500"))
		return
	}
	ten_sach := string(book.Ten_sach)
	tac_gia := string(book.Ten_tg)
	cate := string(book.Cate)
	var idBook int

	selDB, err := db.Prepare("INSERT INTO longphu.book(Ten_sach) VALUES (?) ")
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
		return
	}

	selDB.Exec(ten_sach)

	err1 := db.QueryRow("SELECT book.idBook FROM longphu.book as book WHERE book.Ten_sach = ?", ten_sach).Scan(&idBook)
	if err1 != nil || err1 == sql.ErrNoRows {
		json.NewEncoder(response).Encode(ResponseWriter(err1, "400"))
		return
	}

	author := strings.Split(tac_gia, "; ")

	for i := 0; i < len(author); i++ {
		var idAuthor int
		err1 := db.QueryRow("SELECT author.idAuthor FROM longphu.author as author WHERE author.ten_tg = ?", author[i]).Scan(&idAuthor)
		if err1 != nil || err1 == sql.ErrNoRows {
			json.NewEncoder(response).Encode(ResponseWriter(err1, "400"))
			return
		}

		selDB1, err := db.Prepare("INSERT INTO longphu.book_author (idAuthor, idBook) VALUES (?, ?) ")
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
			return
		}
		selDB1.Exec(idAuthor, idBook)
	}

	categories := strings.Split(cate, "; ")
	for i := 0; i < len(author); i++ {
		var idCategories int
		err1 := db.QueryRow("SELECT categories.idCategories FROM longphu.categories as categories WHERE categories.The_loai = ?", categories[i]).Scan(&idCategories)
		if err1 != nil || err1 == sql.ErrNoRows {
			json.NewEncoder(response).Encode(ResponseWriter(err1, "400"))
			return
		}

		selDB1, err := db.Prepare("INSERT INTO longphu.book_categories (idCategories, idBook) VALUES (?, ?) ")
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
			return
		}
		selDB1.Exec(idCategories, idBook)
	}

	defer selDB.Close()
	res := "INSERT: Ten sach: " + ten_sach + " | Tac gia: " + tac_gia + " | The loai: " + cate
	json.NewEncoder(response).Encode(res)
}

func SearchBookByCate(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "500"))
		return
	}
	bodyString := "%" + string(body) + "%"
	selDB, err := db.Query("SELECT book.Ten_sach, categories.The_loai FROM longphu.book as book,  longphu.categories as categories, longphu.book_categories as bc WHERE book.idBook = bc.idBook AND bc.idCategories = categories.idCategories AND categories.The_loai LIKE ?", bodyString)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
		return
	}
	defer selDB.Close()
	book := Book{}
	res := []Book{}
	for selDB.Next() {
		var Ten_sach string
		var The_loai string
		err = selDB.Scan(&Ten_sach, &The_loai)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
			return
		}
		book.Ten_sach = Ten_sach
		book.Cate = The_loai
		res = append(res, book)
	}
	json.NewEncoder(response).Encode(res)
}

func SearchBookByAuthor(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "500"))
		return
	}
	bodyString := "%" + string(body) + "%"
	selDB, err := db.Query("SELECT book.Ten_sach, author.Ten_Tg FROM longphu.book as book,  longphu.author as author, longphu.book_author as ba WHERE book.idBook = ba.idBook AND ba.idAuthor = author.idAuthor AND author.Ten_Tg LIKE ?", bodyString)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
		return
	}
	defer selDB.Close()
	book := Book{}
	res := []Book{}
	for selDB.Next() {
		var Ten_sach string
		var Ten_tg string
		err = selDB.Scan(&Ten_sach, &Ten_tg)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
			return
		}
		book.Ten_sach = Ten_sach
		book.Ten_tg = Ten_tg
		res = append(res, book)
	}
	json.NewEncoder(response).Encode(res)
}
