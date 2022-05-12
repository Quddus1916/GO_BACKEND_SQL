package models

import (
	"github.com/Quddus1916/GO_BACKEND_SQL/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Id      int8   `gorm:"" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type Id struct {
	gorm.Model
	Id int8 `gorm:"" json:"id"`
}

func init() {
	database.Connect()

	db = database.GetDB()
	//table name is defined through this
	db.AutoMigrate(&User{})
}
func CreateUser(user *User) *User {
	db.NewRecord(user)
	db.Create(&user)

	return user
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(id *Id) (*User, *gorm.DB) {
	var sigleuser User
	db := db.Where("id=?", id.Id).Find(&sigleuser)
	return &sigleuser, db

}

func DeleteUser(id int64) User {
	var user User
	db.Where("id=?", id).Delete(user)
	return user

}
