package models

import (
	"Quddus1916/GO_BACKEND_SQL/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Userdata struct {
	gorm.Model
	U_Id     int8   `gorm:"" json:"id"`
	Uname    string `json:"name"`
	Uaddress string `json:"address"`
}
type Id struct {
	gorm.Model
	Id int8 `gorm:"" json:"id"`
}

func init() {
	database.Connect()

	db = database.GetDB()
	//table name is defined through this
	db.AutoMigrate(&Userdata{})
}
func CreateUser(user *Userdata) *Userdata {
	db.NewRecord(user)
	db.Create(&user)

	return user
}

func GetAllUsers() []Userdata {
	var Users []Userdata
	db.Find(&Users)
	return Users
}

func GetUserById(id *Id) (*Userdata, *gorm.DB) {
	var sigleuser Userdata
	db := db.Where("U_ID=?", id.Id).Find(&sigleuser)
	return &sigleuser, db

}

func DeleteUser(id int64) Userdata {
	var user Userdata
	db.Where("U_ID=?", id).Delete(user)
	return user

}
