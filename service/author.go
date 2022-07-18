package service

import (
	"book-sto/model"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Author = model.Author

type ret map[string]string

func ResponseWriter(err error, status string) map[string]string {
	ret := make(map[string]string)
	ret["status"] = status
	ret[err.Error()] = err.Error()
	return ret
}

func IndexAuthor(response http.ResponseWriter, request *http.Request) {
	selDB, err := db.Query("SELECT idAuthor,ten_tg, QueQuan FROM author")
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
	}
	defer selDB.Close()
	author := Author{}
	res := []Author{}
	for selDB.Next() {
		var idAuthor int
		var ten_tg string
		var QueQuan string
		err = selDB.Scan(&idAuthor, &ten_tg, &QueQuan)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
			return
		}
		author.IdAuthor = idAuthor
		author.Ten_tg = ten_tg
		author.QueQuan = QueQuan
		res = append(res, author)
	}
	json.NewEncoder(response).Encode(res)

}

func SearchAuthor(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
	}
	bodyString := "%" + string(body) + "%"
	selDB, err := db.Query("SELECT author.ten_tg, author.QueQuan FROM longphu.author as author WHERE author.ten_tg LIKE ?", bodyString)
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
		return
	}
	defer selDB.Close()
	author := Author{}
	res := []Author{}
	for selDB.Next() {
		var ten_tg string
		var QueQuan string
		err = selDB.Scan(&ten_tg, &QueQuan)
		if err != nil {
			json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
			return
		}
		author.Ten_tg = ten_tg
		author.QueQuan = QueQuan
		res = append(res, author)
	}
	json.NewEncoder(response).Encode(res)
}

func CreateAuthor(response http.ResponseWriter, request *http.Request) {
	var author Author
	if err := json.NewDecoder(request.Body).Decode(&author); err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
		return
	}
	queQuan := string(author.QueQuan)
	ten_tg := string(author.Ten_tg)

	selDB, err := db.Prepare("INSERT INTO longphu.author(ten_tg, QueQuan) VALUES (? ,?) ")
	if err != nil {
		json.NewEncoder(response).Encode(ResponseWriter(err, "400"))
		return
	}
	selDB.Exec(ten_tg, queQuan)
	defer selDB.Close()
	res := "INSERT: ten tac gia: " + ten_tg + " | Que quan: " + queQuan
	json.NewEncoder(response).Encode(res)
}
