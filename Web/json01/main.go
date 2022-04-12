package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func index(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>FOO</title>
		</head>
		<body>
		You are at foo
		</body>
		</html>`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "DTher",
		Lname: "Htun",
		Items: []string{"laptop", "phone", "backpack"},
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}

func encd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p2 := person{
		"Wai",
		"Yan",
		[]string{"laptop", "phone", "backpack"},
	}
	err := json.NewEncoder(w).Encode(p2)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
