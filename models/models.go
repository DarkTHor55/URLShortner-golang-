package models

import (
	"time"
	"url-shortner/dbConnection"
	"gorm.io/gorm"
)

type URL struct {
	ID          string    `gorm:"primaryKey;type:varchar(255);not null" json:"id"` 
	OriginalURL string    `gorm:"type:text;not null" json:"original_url"`
	CreationURL string    `gorm:"type:varchar(255);not null" json:"creation_url"`
	CurrentDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"current_date"`
}

func init() {
	dbConnection.ConnectDB()
	db := dbConnection.GetDB()

	if err := db.AutoMigrate(&URL{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

// BeforeCreate is called before a new URL is created.
// It sets the ID to be the same as CreationURL.
func (url *URL) BeforeCreate(tx *gorm.DB) (err error) {
	// The ID will be set to the CreationURL value before creating the record
	url.ID = url.CreationURL
	return nil
}

func (url *URL) CreateUrl() (*URL, error) {
	db := dbConnection.GetDB()
	if err := db.Create(url).Error; err != nil { 
		return nil, err 
	}
	return url, nil 
}

func GetAllURL() ([]URL, error) {
	db := dbConnection.GetDB()
	var urls []URL
	if err := db.Find(&urls).Error; err != nil {
		return nil, err 
	}
	return urls, nil
}

// GetURLByID now retrieves the URL based on the CreationURL.
func (url *URL) GetURLByID(shortURL string) (*URL, error) {
	db := dbConnection.GetDB()
	var retrievedURL URL
	result := db.Where("creation_url = ?", shortURL).First(&retrievedURL)
	if result.Error != nil {
		return nil, result.Error
	}
	return &retrievedURL, nil
}
