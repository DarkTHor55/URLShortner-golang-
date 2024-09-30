package main

import (
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"
	"url-shortner/models"
	"url-shortner/dbConnection"

)


var urlDb = make(map[string]URL)
func genrateShortUrl(OrignalURL string)string {
	// md5 (Message-Digest Algorithm 5) is use to convert string into 128-bit hash value
	hash:= md5.New()

	hash.Write([]byte(OrignalURL))// it cconvert string into btye slice

	// hash have all bytes

	data := hash.Sum(nil)//contain all data in slice but in byte

	// convert byte slice to string
	hashString := hex.EncodeToString(data)

	// fmt.Print(hashString)

	return hashString[:8] //we only return 8 char of string 

}

func createUrl(OrignalURL string)(*models.URL, error){
	shortURL := genrateShortUrl(OrignalURL)

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

	// Example original URL
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