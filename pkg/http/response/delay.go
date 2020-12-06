package response

import (
	"fmt"
	"net/http"
	"time"
)

// DelayHandleFunc to be used as http.HandleFunc for delay response by sleeping passed in time, e.g. 5s, 1200ms
func DelayHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	delayPath := r.URL.Path[len("/response/delay/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		var message = "Bad request"
		if delay, err := time.ParseDuration(delayPath); err == nil {
			message = fmt.Sprintf("Delayed by: %s", delay.String())
			time.Sleep(delay)
			w.Header().Add("Response-Delay", delay.String())
			w.WriteHeader(http.StatusOK)
		} else if delayPath != "" {
			message = err.Error()
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write([]byte(message))
		// fmt.Fprintf(w, message)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}

}
