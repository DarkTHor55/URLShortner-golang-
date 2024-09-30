package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
	"url-shortner/models"
	"url-shortner/dbConnection"
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
	dbConnection.ConnectDB()

	fmt.Println("Starting URL shortener...")

	originalURL := "https://chatgpt.com/c/66faa9e0-b018-800c-84d7-e8ba5ac41dba"

	// Create the short URL and save it in the database
	newURL, err := createUrl(originalURL)
	if err != nil {
		fmt.Println("Error creating short URL:", err)
		return
	}

	fmt.Println("Original URL:", newURL.OriginalURL)
	fmt.Println("Generated short URL:", newURL.CreationURL)
}
