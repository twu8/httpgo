package api

import (
	"fmt"
	"net/http"
)

func EchoHandleFunc(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/plain")
	if len(r.URL.Query()["message"]) == 0 {
		fmt.Fprintf(w, "No message to echo")
		return
	}
	message := r.URL.Query()["message"][0]
	fmt.Fprintf(w, message)
}
