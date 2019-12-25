package main

import (
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Setup proxy
	url1, err := url.Parse("http://localhost:8082")
	if err != nil {
		e.Logger.Fatal(err)
	}

	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
	}

	// All requests starting with "stream" will be forwarded to targets
	g := e.Group("/stream")
	g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.Logger.Fatal(e.Start(":8081"))
}