package model

import (
	"log"
	"github.com/taroooth/order/config"
	"fmt"
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

func init() {
	var err error
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)

	db, err = gorm.Open(mysql.Open(dbConnectInfo), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connect database..")
	}
}
