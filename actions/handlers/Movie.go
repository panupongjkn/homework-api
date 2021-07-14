package handlers

import (
	"homework-api/actions/repositories"
	"homework-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMovie() echo.HandlerFunc {
	fn := func(c echo.Context) error {
		all_movie, err := repositories.GetMovie()
		if err != nil {
			message := models.NewMessage(404, "empty movie")
			return c.JSON(http.StatusOK, message)
		}
		return c.JSON(http.StatusOK, all_movie)
	}
	return echo.HandlerFunc(fn)
}
