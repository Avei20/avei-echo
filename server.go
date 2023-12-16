package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	port := os.Getenv("PORT")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World from Avei!")
	})

	e.Logger.Fatal(e.Start(":"+port))
}