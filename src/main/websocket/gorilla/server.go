package main

import (
	"github.com/labstack/echo/middleware"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var(
	upgrader = websocket.Upgrader{}
)

func hello (c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		//write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil{
			c.Logger().Error(err)
		}

		//Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Println("%s\n", msg)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}