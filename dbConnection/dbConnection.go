package dbConnection

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB

const (
	username = "root"
	password = "Rohit@123"
	hostname = "127.0.0.1:3306"
	dbname   = "urlshortner"
)

// Initialize MySQL Connection
func ConnectDB() {
	// MySQL connection string: user:password@tcp(host:port)/dbname
	dsn := ("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)

	// Open the connection
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening the database: ", err)
	}

	// Check if the connection is available
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	fmt.Println("Successfully connected to the MySQL database!")
}
func GetDB() *sql.DB {
	if DB == nil {
		ConnectDB()
	}
	return DB
}