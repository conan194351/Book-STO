package service

import (
	"book-sto/config"
	"book-sto/model"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Categories model.Categories

func IndexCate(response http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	selDB, err := db.Query("SELECT * FROM longphu.categories ")
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}
	defer selDB.Close()
	categories := Categories{}
	res := []Categories{}
	for selDB.Next() {
		var idCategories int
		var The_loai string
		err = selDB.Scan(&idCategories, &The_loai)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err))
		}
		categories.IdCategories = idCategories
		categories.The_loai = The_loai
		res = append(res, categories)
	}
	json.NewEncoder(response).Encode(res)
}

func CreateCate(response http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	var categories Categories
	if err := json.NewDecoder(request.Body).Decode(&categories); err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}
	The_loai := categories.The_loai

	selDB, err := db.Prepare("INSERT INTO longphu.categories(The_loai) VALUES (?) ")
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}
	selDB.Exec(The_loai)
	defer selDB.Close()
	res := "INSERT: The loai: " + The_loai
	json.NewEncoder(response).Encode(res)
}

func SearchCate(response http.ResponseWriter, request *http.Request) {
	db := config.DbConn()
	body, err := ioutil.ReadAll(request.Body)
	bodyString := "%" + string(body) + "%"
	selDB, err := db.Query("SELECT categories.The_loai FROM longphu.categories as categories WHERE categories.The_loai LIKE ?", bodyString)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err))
	}
	defer selDB.Close()
	categories := Categories{}
	res := []Categories{}
	for selDB.Next() {
		var The_loai string
		err = selDB.Scan(&The_loai)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err))
		}
		categories.The_loai = The_loai
		res = append(res, categories)
	}
	json.NewEncoder(response).Encode(res)
}
