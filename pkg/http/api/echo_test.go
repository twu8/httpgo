package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoHandleFunc(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		body           string // For POST requests, though current impl ignores it
		expectedStatus int
		expectedBody   string
		isQueryParam   bool   // True if we are testing query param, false for others
	}{
		{
			name:           "GET with message",
			method:         "GET",
			path:           "/api/echo?message=hello",
			expectedStatus: http.StatusOK,
			expectedBody:   "hello",
			isQueryParam:   true,
		},
		{
			name:           "GET no message",
			method:         "GET",
			path:           "/api/echo",
			expectedStatus: http.StatusOK, // The handler itself doesn't set error status for "No message"
			expectedBody:   "No message to echo",
			isQueryParam:   true,
		},
		{
			name:           "GET empty message",
			method:         "GET",
			path:           "/api/echo?message=",
			expectedStatus: http.StatusOK,
			expectedBody:   "No message to echo", // Based on `len(r.URL.Query()["message"]) == 0`
			isQueryParam:   true,
		},
		{
			name:           "POST with message in query",
			method:         "POST",
			path:           "/api/echo?message=postquery",
			body:           "This is a post body, should be ignored",
			expectedStatus: http.StatusOK,
			expectedBody:   "postquery",
			isQueryParam:   true,
		},
		{
			name:           "POST with no message in query",
			method:         "POST",
			path:           "/api/echo",
			body:           "This is a post body, should be ignored",
			expectedStatus: http.StatusOK,
			expectedBody:   "No message to echo",
			isQueryParam:   true,
		},
		{
			name:           "PUT with message in query (behavior check)",
			method:         "PUT",
			path:           "/api/echo?message=putquery",
			body:           "This is a put body, should be ignored",
			expectedStatus: http.StatusOK,
			expectedBody:   "putquery",
			isQueryParam:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			var err error

			if tt.method == "POST" || tt.method == "PUT" {
				req, err = http.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			} else {
				req, err = http.NewRequest(tt.method, tt.path, nil)
			}
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(EchoHandleFunc)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			responseBody, err := ioutil.ReadAll(rr.Body)
			if err != nil {
				t.Fatalf("could not read response body: %v", err)
			}

			if strings.TrimSpace(string(responseBody)) != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got '%s' want '%s'",
					string(responseBody), tt.expectedBody)
			}

			contentType := rr.Header().Get("Content-Type")
			if contentType != "text/plain" {
				t.Errorf("handler returned wrong content type: got %s want text/plain", contentType)
			}
		})
	}
}
