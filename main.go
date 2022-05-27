package main

import (
	"Quddus1916/GO_BACKEND_SQL/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: " uri=${uri},method=${method}, status=${status}\n",
	}))

	e.GET("/users", controller.GetUsers)
	e.POST("/user", controller.CreateUser)
	e.GET("/user", controller.GetUserById)

	e.Start(":8080")

}
