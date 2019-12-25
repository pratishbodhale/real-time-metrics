package controllers

import (
	"encoding/json"
	"fmt"
	echo "github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
	"stream/models"
)

// Controller interface has two methods
type Controller interface {
	// Homecontroller renders initial home page
	HomeController(e echo.Context) error

	// StreamController responds with live cpu status over websocket
	StreamController(e echo.Context) error
}

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

var model models.Model

// Initializes the models
func Init(){
	model = models.NewModel()
}



func (c *controller) HomeController(e echo.Context) error {
	return e.File("views/index.html")
}

func (c *controller) StreamController(e echo.Context) error {

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		status, err := model.GetLiveCpuUsage()
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			// Write
			newVal := <- status
			jsonResponse, _ := json.Marshal(newVal)
			err := websocket.Message.Send(ws, fmt.Sprintln(string(jsonResponse)))
			if err != nil {
				fmt.Println(err)
			}
		}
	}).ServeHTTP(e.Response(), e.Request())
	return nil
}