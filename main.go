package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/book", ViewBook)
	http.HandleFunc("/books", ViewBooks)
	http.HandleFunc("/books/new", NewBook)
	http.HandleFunc("/books/create", RegisterBook)
	http.HandleFunc("/books/edit", UpDateForm)
	http.HandleFunc("/books/update", BookUpdate)
	http.HandleFunc("/books/delete", DeleteBook)
	http.HandleFunc("/books/confirmdelete", ConfirmDelete)

	fmt.Println("server runnig")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
