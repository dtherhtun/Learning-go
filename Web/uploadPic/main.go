package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCooike(w, r)

	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		mf.Seek(0, 0)
		io.Copy(nf, mf)
		c = appenCookie(w, c, fname)
	}
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.html", xs)
}

func getCooike(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: uuid.NewString(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appenCookie(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
