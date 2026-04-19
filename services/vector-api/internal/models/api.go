package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type Error struct {
	Code    int
	Message string
}

// Error handling
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

// Creating specific instances for custom use
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	// Internal server error does not need an external error code
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)

// Collection handler schemas
type CreateDeleteCollectionRequest struct {
	Name string `json:"name"`
}

type CreateDeleteCollectionResponse struct {
	Code             int
	Name             string    `json:"name"`
	CreatedDeletedAt time.Time `json:"created_delete_at"`
}

type RenameCollectionRequest struct {
	Name    string `json:"name"`
	NewName string `json:"new_name"`
}

type RenameCollectionResponse struct {
	Code      int
	NewName   string    `json:"new_name"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Entry handler schemas
type AddNewEntryRequest struct {
	ID     []int64  `json:"id"`
	Name   []string `json:"name"`
	Album  []string `json:"album"`
	Artist []string `json:"artist"`
	Mood   []string `json:"mood"`
}

type AddNewEntryResponse struct {
	Code       int
	InsertedAt time.Time `json:"inserted_at"`
}

type UpdateEntryRequest struct {
	ID     []int64  `json:"id"`
	Name   []string `json:"name"`
	Album  []string `json:"album"`
	Artist []string `json:"artist"`
	Mood   []string `json:"mood"`
}

type QueryEntryRequest struct {
	Query string `json:"query"` // This should be a link to an image?
}

type QueryEntryResponse struct {
	Code             int
	RecommendedSongs []string `json:"recommended_songs"` // Should return the list of titles
}

type DeleteEntryRequest struct {
	ID []int64 `json:"id"`
}

type DeleteEntryResponse struct {
	Code int
}
