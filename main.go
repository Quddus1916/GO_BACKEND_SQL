package main

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

)

func main (){
	e:= echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: " uri=${uri},method=${method}, status=${status}\n",
	  }))
	 

	e.GET("/users",controller.Users)
	e.POST("/users",controller.CreateUser)
	
	e.Start(":8080")
	 
	  
}