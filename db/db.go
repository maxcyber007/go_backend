package db
import (
	"os"
	"log"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
	var Db *gorm.DB
	var err error

	func InitDB() { 
	dns := os.Getenv("POSTGRES_DNS") 
	Db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil { 
		log.Fatal("Failed to get database connection:", err) 
	} 
	fmt.Println("Connect Database Successfuly!")
	} 