package routes

import (
	"github.com/labstack/echo/v4"
	"stream/controllers"
)

func Init(e *echo.Echo){

	controllers.Init()
	c := controllers.NewController()
	
	e.GET("/", c.HomeController)
	e.GET("/stream", c.StreamController)
}
