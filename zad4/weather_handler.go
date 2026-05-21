package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

const URL = "https://danepubliczne.imgw.pl/api/data/synop/station/"

type WeatherHandler struct {
	DB *gorm.DB
}

func (h *WeatherHandler) GetWeather(c *echo.Context) error {
	stationsParam := c.QueryParam("stations")

	if stationsParam == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "stations parameter is required",
		})
	}

	stations := strings.Split(stationsParam, ",")

	resultJSON := make([]Weather, 0)

	for _, station := range stations {

		proxy := WeatherProxy{URL: URL}

		weather, err := proxy.FetchStation(station)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		resultJSON = append(resultJSON, weather)
	}

	// Save the results to the database
	h.DB.Create(&resultJSON)

	return c.JSON(http.StatusOK, resultJSON)
}
