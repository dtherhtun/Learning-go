package customer

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB() *sql.DB {
	os.Remove("./customer.db")

	log.Println("Creating customer.db...")
	file, err := os.Create("./customer.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("customer.db created")

	sqlDatabase, _ := sql.Open("sqlite3", "./customer.db")
	defer sqlDatabase.Close()
	createTable(sqlDatabase)

	return sqlDatabase
}

func createTable(db *sql.DB) {
	log.Println("Creating table...")

	_, err := db.Exec("CREATE TABLE customers (id integer NOT NULL PRIMARY KEY AUTOINCREMENT,	username TEXT NOT NULL, password TEXT NOT NULL, email TEXT NOT NULL	)")

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Table created")
}
