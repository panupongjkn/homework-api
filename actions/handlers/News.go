package handlers

import (
	"homework-api/actions/repositories"
	"homework-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetNews() echo.HandlerFunc {
	fn := func(c echo.Context) error {
		all_news, err := repositories.GetNews()
		if err != nil {
			message := models.NewMessage(404, "empty news")
			return c.JSON(http.StatusOK, message)
		}
		return c.JSON(http.StatusOK, all_news)
	}
	return echo.HandlerFunc(fn)
}

func GetNewsSport() echo.HandlerFunc {
	fn := func(c echo.Context) error {
		all_news, err := repositories.GetNewsSport()
		if err != nil {
			message := models.NewMessage(404, "empty news")
			return c.JSON(http.StatusOK, message)
		}
		return c.JSON(http.StatusOK, all_news)
	}
	return echo.HandlerFunc(fn)
}
