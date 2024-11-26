package db
import (
	_"os"
	"log"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
	var Db *gorm.DB
	var err error

	func InitDB() { 
	//dsn := os.Getenv(POSTGRES_DNS) 
	POSTGRES_DNS := "postgres://root:p@sswd@2024@itdev.cmtc.ac.th:5432/db_67319010001"
	Db, err = gorm.Open(postgres.Open(POSTGRES_DNS), &gorm.Config{})
	if err != nil { 
		log.Fatal("Failed to get database connection:", err) 
	} 
	fmt.Println("Connect Database Successfuly!")
	} 