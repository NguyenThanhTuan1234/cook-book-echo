package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/request", func(c echo.Context) error{
		req := c.Request()
		format := `
			<code>
				Protocol: %s<br>
				Host: %s<br>
				Remote Address: %s<br>
				Method: %s<br>
				Path: %s<br>
			</code>
		`
		return c.HTML(http.StatusOK,fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))

	})
	e.Logger.Fatal(e.StartTLS("8000", "cert.pem", "key.pem"))
}