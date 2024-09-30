package controller

import (
	"encoding/json"
	"net/http"
	"url-shortner/models"
	"github.com/gorilla/mux" 
)

// CreateURL handles the creation of a new short URL
func CreateURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newURL models.URL
	if err := json.NewDecoder(r.Body).Decode(&newURL); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdURL, err := newURL.CreateUrl()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdURL)
}

// GetAllURLs retrieves all stored short URLs
func GetAllURLs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urls, err := models.GetAllURL() 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(urls)
}

// GetURLByID retrieves a short URL by its ID
func GetURLByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var url models.URL
	result, err := url.GetURLByID(params["id"]) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
}
