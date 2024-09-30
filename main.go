package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
	"net/http"
	"url-shortner/models"
	"url-shortner/dbConnection"
	"url-shortner/router"
)

func generateShortUrl(originalURL string) string {
	// md5 (Message-Digest Algorithm 5) is used to convert string into a 128-bit hash value
	hash := md5.New()

	hash.Write([]byte(originalURL)) // Convert string into byte slice

	// Hash contains all bytes
	data := hash.Sum(nil) // Contains all data in slice but in byte

	// Convert byte slice to string
	hashString := hex.EncodeToString(data)

	return hashString[:8] // Return only the first 8 characters of the hash
}

func createUrl(originalURL string) (*models.URL, error) {
	shortURL := generateShortUrl(originalURL)

	newURL := &models.URL{
		ID:shortURL,
		OriginalURL: originalURL,
		CreationURL: shortURL,
		CurrentDate: time.Now(),
	}

	if _, err := newURL.CreateUrl(); err != nil {
		return nil, err // Return an error if the insertion fails
	}

	return newURL, nil
}


func main() {
	// Initialize database connection
	dbConnection.ConnectDB() // No need for an error check since it doesn't return a value

	fmt.Println("Starting URL shortener...")

	// Generate a short URL for an example original URL
	originalURL := "https://chatgpt.com/c/66faa9e0-b018800c-84d7-e8ba5ac41dba"

	// Create and store the short URL in the database
	newURL, err := createUrl(originalURL)
	if err != nil {
		fmt.Println("Error creating short URL:", err)
		return
	}

	// Output the original and short URLs
	fmt.Println("Original URL:", newURL.OriginalURL)
	fmt.Println("Generated short URL:", newURL.CreationURL)

	// Setup router and start the HTTP server
	r := router.SetupRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}