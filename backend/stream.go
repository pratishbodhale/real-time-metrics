package stream

import (
	"github.com/labstack/echo/v4"
	"stream/routes"
)

func Init(){
	e := echo.New()

	// registers all the available routes
	routes.Init(e)

	// starts echo server
	e.Logger.Fatal(e.Start(":8082"))
}