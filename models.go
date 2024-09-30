// package models

// import (
// 	"fmt"
// 	"project/db" // Import the database package for DB connection
// 	"time"
// )

// // Define your struct
// type URL struct {
// 	ID          string    `json:"id"`
// 	OrignalURL  string    `json:"orignal_url"`
// 	CreationURL string    `json:"creation_url"`
// 	CurrentDate time.Time `json:"current_date"`
// }
// func InsertURL(url URL) error {
// 	query := "INSERT INTO urls (id, orignal_url, creation_url, current_date) VALUES (?, ?, ?, ?)"
// 	_, err := db.DB.Exec(query, url.ID, url.OrignalURL, url.CreationURL, url.CurrentDate)
// 	if err != nil {
// 		return fmt.Errorf("failed to insert URL: ", err)
// 	}
// 	return nil
// }

// func GetURLByID(id string) (URL, error) {
// 	var url URL
// 	query := "SELECT id, orignal_url, creation_url, current_date FROM urls WHERE id = ?"
// 	err := db.DB.QueryRow(query, id).Scan(&url.ID, &url.OrignalURL, &url.CreationURL, &url.CurrentDate)
// 	if err != nil {
// 		return URL{}, fmt.Errorf("failed to get URL: ", err)
// 	}
// 	return url, nil
// }