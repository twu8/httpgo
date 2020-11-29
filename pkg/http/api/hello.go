package api

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Message string
}

func HelloHandleFunc(w http.ResponseWriter, r *http.Request) {

	m := Hello{"Welcome to Cloud Native Go."}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
