package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")

	//JSONP
	e.GET("/jsonp", func(c echo.Context) error {
		callback := c.QueryParam("callback")
		var content struct {
			Response	string		`json :"response"`
			Timestamp	time.Time	`json:"timestamp"`
			Random 		int			`json:"random"`
		}
		content.Response = "Sent via JSONP"
		content.Timestamp = time.Now().UTC()
		content.Random = rand.Intn(1000)
		return c.JSONP(http.StatusOK, callback, &content)
	})
	e.Logger.Fatal(e.Start(":8000"))
}