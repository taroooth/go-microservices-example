package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	Id       uint64 `json:"id,omitempty"`
	AuthorId uint64 `json:"author_id,omitempty"`
	Title    string `gorm:"size:255" json:"title,omitempty"`
}

func GetAllBooks(books *[]Book) {
	db.Find(&books)
}

func CreateBook(book *Book) {
	db.Create(&book)
}

func UpdateBook(book *Book) {
	db.Save(&book)
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Failed to load env file: %v", err)
		log.Fatalln(err)
	}

	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = gorm.Open(mysql.Open(dbConnectInfo), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connect database..")
	}
}
