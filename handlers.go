package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func renderTemplate(w http.ResponseWriter, file string, data any) {
	tmpl, err := template.ParseFiles("templates/base.html", file)
	if err != nil {
		http.Error(w, "Cannot parse file", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, "templates/home.html", nil)
}

func ViewBook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/book" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")

	if id == "" {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	convert_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID not found", http.StatusBadRequest)
		return
	}

	for _, book := range books {
		if book.ID == convert_id {
			renderTemplate(w, "templates/book.html", book)
			return
		}
	}
	http. Error(w, "Book not found", http.StatusBadRequest)
	return
}

func ViewBooks(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/books" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w,"Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	renderTemplate(w, "templates/books.html", books)
	return
}

func NewBook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/books/new" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	renderTemplate(w, "templates/newbook.html", nil)
}

func RegisterBook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/books/create" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	title := r.FormValue("title")
	author := r.FormValue("author")
	publisheddate := r.FormValue("publisheddate")
	edition := r.FormValue("edition")

	if id == "" || author == "" || title == "" || publisheddate == "" || edition == "" {
		http.Error(w,"One or more parameter missing", http.StatusBadRequest)
		return
	}

	convert_publisheddate, err := strconv.Atoi(publisheddate)
	if err != nil {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	convert_edition, err := strconv.Atoi(edition)
	if err != nil {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	nextid := len(books) + 1

	newbook := Books{
		ID: nextid,
		Title: title,
		Author: author,
		PublishedDate: convert_publisheddate,
		Edition: convert_edition,
	}
	books = append(books, newbook)
	http.Redirect(w,r,"/books", http.StatusSeeOther)
}

func UpDateForm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/books/edit" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")

	if id == "" {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	convert_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	for _,book := range books {
		if book.ID == convert_id {
			renderTemplate(w, "templates/updateform.html",book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusBadRequest)
	return
}

func BookUpdate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/books/update" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	title := r.FormValue("title")
	author := r.FormValue("author")
	publisheddate := r.FormValue("publisheddate")
	edition := r.FormValue("edition")

	if id == "" || publisheddate == "" || author == "" || title == "" || edition == "" {
		http.Error(w, "One or more parameter missing", http.StatusBadRequest)
		return
	}

	convert_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}

	convert_publisheddate, err := strconv.Atoi(publisheddate)
	if err != nil {
		http.Error(w, "Publish date required", http.StatusBadRequest)
		return
	}

	convert_edition, err := strconv.Atoi(edition)
	if err != nil {
		http.Error(w, "Edition required", http.StatusBadRequest)
		return
	}

	for i := range books {
		if books[i].ID == convert_id {
			books[i].Title = title
			books[i].Author = author
			books[i].PublishedDate = convert_publisheddate
			books[i].Edition = convert_edition
			http.Redirect(w,r,"/books", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusBadRequest)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/books/delete" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")

	convert_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	for i := range books {
		if books[i].ID == convert_id {
			books = append(books[:i], books[i + 1:]...)
			http.Redirect(w, r, "/books", http.StatusSeeOther)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusBadRequest)
}

func ConfirmDelete(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/confirmdelete" {
		http.NotFound(w,r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")

	convert_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	for _,book := range books {
		if book.ID == convert_id {
			renderTemplate(w, "templates/confirmdelete", book)
			return
		}
	}
	http.Error(w, "Book not Found", http.StatusBadRequest)
}