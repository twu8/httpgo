package response

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"sort"
)

// Helper function to sort semicolon-separated key-value pairs for consistent comparison
func sortHeaderString(s string) string {
	if s == "" {
		return ""
	}
	parts := strings.Split(s, "; ")
	sort.Strings(parts)
	return strings.Join(parts, "; ")
}

func TestHeaderHandleFunc(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		headers        map[string]string // Simplified: single value per header for testing
		multiValueHeaders map[string][]string // For headers with multiple values
		expectedStatus int
		expectedBody   string // Expected body parts, will be sorted for comparison
	}{
		{
			name:           "GET no headers",
			method:         "GET",
			headers:        nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
		{
			name:           "GET single header",
			method:         "GET",
			headers:        map[string]string{"X-Test-Header": "Value1"},
			expectedStatus: http.StatusOK,
			expectedBody:   "x-test-header=Value1",
		},
		{
			name:           "GET multiple headers",
			method:         "GET",
			headers:        map[string]string{"X-Test-Header-A": "ValueA", "X-Test-Header-B": "ValueB"},
			expectedStatus: http.StatusOK,
			expectedBody:   "x-test-header-a=ValueA; x-test-header-b=ValueB",
		},
		{
			name:           "GET header with mixed case",
			method:         "GET",
			headers:        map[string]string{"Cache-Control": "no-cache"},
			expectedStatus: http.StatusOK,
			expectedBody:   "cache-control=no-cache",
		},
		{
			name:   "GET with multiple values for one header",
			method: "GET",
			multiValueHeaders: map[string][]string{
				"X-Multi-Value": {"val1", "val2"},
				"X-Another":     {"anotherVal"},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "x-another=anotherVal; x-multi-value=val1; x-multi-value=val2",
		},
		{
			name:           "POST with headers (method agnostic check)",
			method:         "POST",
			headers:        map[string]string{"X-Post-Header": "PostValue"},
			expectedStatus: http.StatusOK,
			expectedBody:   "x-post-header=PostValue",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/response/header", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			if tt.headers != nil {
				for k, v := range tt.headers {
					req.Header.Set(k, v)
				}
			}
			if tt.multiValueHeaders != nil {
				for k, values := range tt.multiValueHeaders {
					for _, v := range values {
						req.Header.Add(k, v)
					}
				}
			}


			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HeaderHandleFunc)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			responseBodyBytes, err := ioutil.ReadAll(rr.Body)
			if err != nil {
				t.Fatalf("could not read response body: %v", err)
			}
			actualBody := strings.TrimSpace(string(responseBodyBytes))

			// Sort both expected and actual bodies for comparison due to map iteration order
			sortedActualBody := sortHeaderString(actualBody)
			sortedExpectedBody := sortHeaderString(tt.expectedBody)

			if sortedActualBody != sortedExpectedBody {
				t.Errorf("handler returned unexpected body: got '%s' want '%s'",
					sortedActualBody, sortedExpectedBody)
			}

			contentType := rr.Header().Get("Content-Type")
			if contentType != "text/plain" {
				t.Errorf("handler returned wrong content type: got '%s' want 'text/plain'", contentType)
			}
		})
	}
}
