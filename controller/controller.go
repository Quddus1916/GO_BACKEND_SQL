package controller

import(
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	"github.com/Quddus1916/GO_BACKEND_SQL/database"
)

type User struct{
	gorm.Model
	Id      int8  `gorm:"" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var db *gorm.DB

func init(){
	database.Connect()
	db = database.GetDB()

}

func CreateUser(c echo.Context) error{

}