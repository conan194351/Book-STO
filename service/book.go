package service

import (
	"book-sto/config"
	"book-sto/model"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Book model.Book

func IndexBook(response http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	selDB, err := db.Query("SELECT book.idBook,book.Ten_sach,author.ten_tg, categories.The_loai FROM longphu.book as book,longphu.author as author, longphu.categories as categories WHERE author.idAuthor = book.idAuthor AND book.idCategories = categories.idCategories")
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}
	defer selDB.Close()
	book := Book{}
	res := []Book{}
	for selDB.Next() {
		var idBook int
		var ten_sach string
		var ten_tg string
		var cate string
		err = selDB.Scan(&idBook, &ten_sach, &ten_tg, &cate)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err))
		}
		book.IdBook = idBook
		book.Ten_sach = ten_sach
		book.Ten_tg = ten_tg
		book.Cate = cate
		res = append(res, book)
	}
	json.NewEncoder(response).Encode(res)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	var book Book
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}
	ten_sach := string(book.Ten_sach)
	tac_gia := string(book.Ten_tg)
	cate := string(book.Cate)
	var idAuthor int
	var idCate int

	err1 := db.QueryRow("SELECT author.idAuthor FROM longphu.author as author WHERE author.ten_tg = ?", tac_gia).Scan(&idAuthor)
	if err1 != nil || err1 == sql.ErrNoRows {
		json.NewEncoder(response).Encode(ResponseWriter(err1))
	}

	err2 := db.QueryRow("SELECT categories.idCategories FROM longphu.categories as categories WHERE categories.The_loai = ?", cate).Scan(&idCate)
	if err2 != nil || err2 == sql.ErrNoRows {
		json.NewEncoder(response).Encode(ResponseWriter(err2))
	}

	selDB, err := db.Prepare("INSERT INTO longphu.book(Ten_sach, idAuthor, idCategories) VALUES (?, ?, ?) ")
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}

	selDB.Exec(ten_sach, idAuthor, idCate)
	defer selDB.Close()
	res := "INSERT: Ten sach: " + ten_sach + " | Tac gia: " + tac_gia + " | The loai: " + cate
	json.NewEncoder(response).Encode(res)
}

func Show(response http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	idBook := mux.Vars(request)["id"]
	selDB, err := db.Query("SELECT book.Ten_sach, author.Ten_Tg, categories.The_loai FROM longphu.book as book, longphu.categories as categories, longphu.author as author WHERE book.idCategories = categories.idCategories AND book.idAuthor = author.idAuthor AND book.idBook = ? ", idBook)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}
	defer selDB.Close()
	book := Book{}
	for selDB.Next() {
		var Ten_sach string
		var Ten_tg string
		var The_loai string
		err = selDB.Scan(&Ten_sach, &Ten_tg, &The_loai)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err))
		}
		book.Ten_sach = Ten_sach
		book.Ten_tg = Ten_tg
		book.Cate = The_loai
	}
	json.NewEncoder(response).Encode(book)
}

func SearchBookByCate(response http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	body, err := ioutil.ReadAll(request.Body)
	bodyString := "%" + string(body) + "%"
	selDB, err := db.Query("SELECT book.Ten_sach, author.Ten_Tg FROM longphu.book as book, longphu.categories as categories, longphu.author as author WHERE book.idCategories = categories.idCategories AND book.idAuthor = author.idAuthor AND categories.The_loai LIKE ?", bodyString)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}
	defer selDB.Close()
	book := Book{}
	res := []Book{}
	for selDB.Next() {
		var Ten_sach string
		var Ten_tg string
		err = selDB.Scan(&Ten_sach, &Ten_tg)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err))
		}
		book.Ten_sach = Ten_sach
		book.Ten_tg = Ten_tg
		res = append(res, book)
	}
	json.NewEncoder(response).Encode(res)
}

func SearchBookByAuthor(response http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	body, err := ioutil.ReadAll(request.Body)
	bodyString := "%" + string(body) + "%"
	selDB, err := db.Query("SELECT DISTINCT book.Ten_sach, author.Ten_Tg FROM longphu.book as book, longphu.categories as categories, longphu.author as author WHERE book.idCategories = categories.idCategories AND book.idAuthor = author.idAuthor AND author.Ten_Tg LIKE ?", bodyString)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}
	defer selDB.Close()
	book := Book{}
	res := []Book{}
	for selDB.Next() {
		var Ten_sach string
		var Ten_tg string
		err = selDB.Scan(&Ten_sach, &Ten_tg)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err))
		}
		book.Ten_sach = Ten_sach
		book.Ten_tg = Ten_tg
		res = append(res, book)
	}
	json.NewEncoder(response).Encode(res)
}
