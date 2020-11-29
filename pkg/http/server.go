package http

import (
	"fmt"
	"net/http"
	"os"

	"httpgo/pkg/http/api"
)

func Execute(port string) {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandleFunc)
	http.HandleFunc("/api/hello", api.HelloHandleFunc)

	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	http.ListenAndServe(findPort(port), nil)
}

func findPort(port string) string {
	p := os.Getenv("PORT")
	if len(p) == 0 {
		p = port
	}
	return ":" + p
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome httpgo!")
}
