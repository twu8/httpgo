package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionHandleFunc(t *testing.T) {
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(VersionHandleFunc)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := VersionInfo{
		Version:   "dev",
		BuildTime: "unknown",
		CommitHash: "unknown",
	}
	var actual VersionInfo
	if err := json.NewDecoder(rr.Body).Decode(&actual); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if actual.Version != expected.Version {
		t.Errorf("handler returned unexpected version: got %v want %v",
			actual.Version, expected.Version)
	}
	if actual.BuildTime != expected.BuildTime {
		t.Errorf("handler returned unexpected build_time: got %v want %v",
			actual.BuildTime, expected.BuildTime)
	}
	if actual.CommitHash != expected.CommitHash {
		t.Errorf("handler returned unexpected commit_hash: got %v want %v",
			actual.CommitHash, expected.CommitHash)
	}
}
