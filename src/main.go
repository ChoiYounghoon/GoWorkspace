package main

import (
	"net/http"

	"myapp"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":3000"), myapp.NewHandler())
}
