package main

import (
	"cook-book/src/main/twitter/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			//Skip authentication for and signup login request
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	//Database connection
	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	//Create indicates
	if err = db.Copy().DB("twitter").C("users").EnsureIndex(mgo.Index{
		Key: []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	//Initalize handler
	h := &handler.Handler{DB: db}

	//Route
	e.POST("/signup", h.Signup)
	e.POST("login", h.Login)
	e.POST("follow/:id", h.Follow)
	e.POST("/posts", h.CreatePost)
	e.GET("/feed", h.FetchPost)

	//Start server
	e.Logger.Fatal(e.Start(":1323"))
}