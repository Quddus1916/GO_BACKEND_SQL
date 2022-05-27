package database

import (
	//"github.com/go-sql-driver/mysql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	database, err := gorm.Open("mysql", "nafiul:quddus1916@tcp(mysql:3306)/testdatabase?")

	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	} else {
		fmt.Println("connected to db")
	}
	db = database
}

func GetDB() *gorm.DB {

	return db
}
