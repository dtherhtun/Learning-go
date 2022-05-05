package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://gopher:gopherpass@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You Connected to your database.")
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/books", bookIndex).Methods("GET")
	r.HandleFunc("/books/{isbn}", booksShow).Methods("GET")
	r.HandleFunc("/books/create/", booksCreateForm).Methods("GET")
	r.HandleFunc("/books/create", booksCreateProcess).Methods("POST")
	r.HandleFunc("/books/update/{isbn}", booksUpdateForm).Methods("GET")
	r.HandleFunc("/books/update", booksUpdateProcess).Methods("POST")
	r.HandleFunc("/books/delete/{isbn}", booksDeleteProcess).Methods("GET")
	r.Handle("/favicon.icon", r.NotFoundHandler)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func bookIndex(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		bks = append(bks, bk)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "books.html", bks)
}

func booksShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["isbn"] == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", vars["isbn"])

	bk := Book{}
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "show.html", bk)
}

func booksCreateForm(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create.html", nil)
}

func booksCreateProcess(w http.ResponseWriter, r *http.Request) {
	bk := Book{}

	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(w, "Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float32(f64)
	_, err = db.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "created.html", bk)
}

func booksUpdateForm(w http.ResponseWriter, r *http.Request) {
	row := db.QueryRow("SELECT * FROM books WHERE isbn=$1", mux.Vars(r)["isbn"])
	bk := Book{}
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, "Internal Server Error!", http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "update.html", bk)
}

func booksUpdateProcess(w http.ResponseWriter, r *http.Request) {
	bk := Book{}

	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(w, "Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float32(f64)
	_, err = db.Exec("UPDATE books SET isbn=$1, title=$2, author=$3, price=$4 WHERE isbn=$1", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		http.Error(w, "Can't update! Something worng.", http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "updated.html", bk)
}

func booksDeleteProcess(w http.ResponseWriter, r *http.Request) {
	if mux.Vars(r)["isbn"] == "" {
		http.Error(w, "isbn value not found!", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("DELETE from books WHERE isbn=$1", mux.Vars(r)["isbn"])
	if err != nil {
		http.Error(w, "Can't delete it!", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
