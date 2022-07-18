package routes

import (
	"book-sto/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Index() http.Handler {
	routes := mux.NewRouter()

	//REST API routes book
	routes.HandleFunc("/api/book/showall", service.IndexBook).Methods("GET")
	routes.HandleFunc("/api/book/create", service.CreateBook).Methods("POST")
	routes.HandleFunc("/api/book/search-by-author", service.SearchBookByAuthor).Methods("POSt")
	routes.HandleFunc("/api/book/search-by-categories", service.SearchBookByCate).Methods("POSt")

	//REST API routes author
	routes.HandleFunc("/api/author", service.IndexAuthor).Methods("GET")
	routes.HandleFunc("/api/author/create", service.CreateAuthor).Methods("POST")
	routes.HandleFunc("/api/author/search", service.SearchAuthor).Methods("POST")

	//REST API routes categories
	routes.HandleFunc("/api/categories", service.IndexCate).Methods("GET")
	routes.HandleFunc("/api/categories/search", service.SearchCate).Methods("POST")
	routes.HandleFunc("/api/categories/create", service.CreateCate).Methods("POST")
	return routes
}
