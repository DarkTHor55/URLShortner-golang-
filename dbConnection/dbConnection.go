package dbConnection

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	username = "root"
	password = "Rohit@123"
	hostname = "127.0.0.1:3306"
	dbname   = "urlshortner"
)

// Initialize MySQL Connection
func ConnectDB() {
	// MySQL connection string
	dsn := username + ":" + password + "@tcp(" + hostname + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

	// Open the connection
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening the database: ", err)
	}

	// Confirm connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error getting database: ", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println("Successfully connected to the MySQL database!")
}

func GetDB() *gorm.DB {
	if DB == nil {
		ConnectDB()
	}
	return DB
}
