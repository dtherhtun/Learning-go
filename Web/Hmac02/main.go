package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("dtherkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		c.Value = e + `|` + getCode(e)
	}

	http.SetCookie(w, c)

	io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	    <a href="/auth">Validate This `+c.Value+`</a>
	  </body>
	</html>`)
}

func auth(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeRcvd := xs[1]
	codeCheck := getCode(email)

	if codeRcvd != codeCheck {
		fmt.Println("HMAC codes didn't mach")
		fmt.Println(codeRcvd)
		fmt.Println(codeCheck)
		http.Redirect(w, r, "/ ", http.StatusSeeOther)
		return
	}

	io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	  	<h1>`+codeRcvd+` - RECEIVED </h1>
	  	<h1>`+codeCheck+` - RECALCULATED </h1>
	  </body>
	</html>`)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/auth", auth)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
