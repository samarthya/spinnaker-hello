// Package main that will launch into the command line
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
 Hellow = "Hello, Docker! <3"
)
func init() {
	log.Println(" initializing the application FN:init ")
}

func rootHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, Hellow)
}

func pingHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
}


// main function
func main() {
	e := echo.New()

	/**
	 * Middleware is a function chained in the HTTP request-response cycle with access to
	 * Echo#Context which it uses to perform a specific action, for example, logging every
	 * request or limiting the number of requests.
	 */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", rootHandler)

	e.GET("/ping", pingHandler)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
