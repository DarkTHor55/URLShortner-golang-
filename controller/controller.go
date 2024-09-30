package controller

import (
	"encoding/json"
	"net/http"
	"url-shortner/models"
)
var url =models.URL

func CreateURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

    var newURL models.URL
    
}