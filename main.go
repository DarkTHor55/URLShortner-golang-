package main

import (
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"
)
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
type URL struct{
	ID string `json:"id"`
	OrignalURL string `json:"orignal_url"`
	CreationURL string `json:"creation_url"`
	CurrentDate time.Time `json:"current_date"`
}

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

func createUrl(OrignalURL string){
	shortURL := genrateShortUrl(OrignalURL)
	
}

func main() {
	fmt.Println("Starting url shortener...")
	str :="https://chatgpt.com/c/66faa9e0-b018-800c-84d7-e8ba5ac41dba"
	shortURL:=genrateShortUrl(str)
	fmt.Println("Generated short URL:", shortURL)
}