package api

import (
	"encoding/json"
	"net/http"
)

// Variables to be set by linker flags
var (
	Version   string = "dev" // Default value
	BuildTime string = "unknown" // Default value
	CommitHash string = "unknown" // Default value
)

type VersionInfo struct {
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
	CommitHash string `json:"commit_hash"`
}

func VersionHandleFunc(w http.ResponseWriter, r *http.Request) {
	info := VersionInfo{
		Version:   Version,
		BuildTime: BuildTime,
		CommitHash: CommitHash,
	}

	b, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
