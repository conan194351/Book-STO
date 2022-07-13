package main

import (
	"book-sto/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routes.Index()
	fmt.Println("Server started with port 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
