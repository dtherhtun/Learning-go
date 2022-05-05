package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

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
}

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(420), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
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

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

func main() {
	http.HandleFunc("/books", index)
	http.HandleFunc("/books/show", booksShow)
	http.Handle("/favicon.icon", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(420), http.StatusMethodNotAllowed)
		return
	}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	bk := Book{}
	err := row.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s, %s, %s, $%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
}
