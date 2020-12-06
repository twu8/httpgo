package response

import (
	"fmt"
	"net/http"
	"strings"
)

// HeaderHandleFunc echos back all the request headers in response body
func HeaderHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	message := ""
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {

			if message == "" {
				message = fmt.Sprintf("%v=%v", name, h)
			} else {
				message = fmt.Sprintf("%v; %v=%v", message, name, h)
			}

		}
	}
	fmt.Fprintf(w, message)
}
