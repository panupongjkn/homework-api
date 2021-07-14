package main

import (
	"fmt"
	"homework-api/actions/handlers"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getPort() string {
	var port = os.Getenv("PORT") // ----> (A)
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port // ----> (B)
}

func main() {
	e := echo.New()
	api := e.Group("/api")
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Homework API")
	})
	api.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Homework API")
	})
	api.GET("/news", handlers.GetNews())
	api.GET("/movies", handlers.GetMovie())
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(getPort()))
}
