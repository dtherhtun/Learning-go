package main

import (
	"context"
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "DTher")

	result := dbAcess(ctx)
	fmt.Fprintln(w, result)
}

func dbAcess(ctx context.Context) int {
	uid := ctx.Value("userID").(int)
	return uid
}

func main() {
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
