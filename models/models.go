package models

import (
	"time"
	"url-shortner/dbConnection"
	"gorm.io/gorm" 
)

type URL struct {
	ID          string    `gorm:"primaryKey;type:char(36);not null;default:UUID()" json:"id"` 
	OriginalURL string    `gorm:"type:text;not null" json:"original_url"`
	CreationURL string    `gorm:"type:varchar(255);not null" json:"creation_url"`
	CurrentDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"current_date"`
}

func init() {
	dbConnection.ConnectDB()
	db := dbConnection.GetDB()

	// Automatically migrate the schema
	if err := db.AutoMigrate(&URL{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

// Automatically generate a new UUID for the ID before creating the record
func (url *URL) BeforeCreate(tx *gorm.DB) (err error) {
	url.ID = uuid.New().String() // Generate a new UUID
	return
}

// CreateUrl inserts a new URL into the database
func (url *URL) CreateUrl() (*URL, error) {
	db := dbConnection.GetDB()
	if err := db.Create(url).Error; err != nil { 
		return nil, err 
	}

	return url, nil 
}

// GetAllURL retrieves all URLs from the database
func GetAllURL() ([]URL, error) {
	db := dbConnection.GetDB()
	var urls []URL
	if err := db.Find(&urls).Error; err != nil {
		return nil, err =
	}
	return urls, nil
}

// GetURLByID retrieves a URL by its ID
func (url *URL) GetURLByID(id string) (*URL, error) {
	db := dbConnection.GetDB()
	result := db.Where("id = ?", id).First(url)
	if result.Error != nil {
		return nil, result.Error
	}
	return url, nil
}
