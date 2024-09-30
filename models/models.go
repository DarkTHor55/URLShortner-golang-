package models

import (
	"time"
	"url-shortner/dbConnection"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type URL struct {
	ID          string    `gorm:"primaryKey;type:char(36);not null" json:"id"` 
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

func (url *URL) BeforeCreate(tx *gorm.DB) (err error) {
	url.ID = uuid.New().String()
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

func (url *URL) GetURLByID(id string) (*URL, error) {
	db := dbConnection.GetDB()
	var retrievedURL URL
	result := db.Where("id = ?", id).First(&retrievedURL)
	if result.Error != nil {
		return nil, result.Error
	}
	return &retrievedURL, nil
}
