package main

import (
	"github.com/glebarez/sqlite"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

func main() {

	e := echo.New()

	e.Use(middleware.RequestLogger())

	db, err := Connect()
	if err != nil {
		e.Logger.Error("failed to connect to database", "error", err)
	}

	db.AutoMigrate(&Weather{})

	weatherHandler := &WeatherHandler{DB: db}
	e.GET("/weather", weatherHandler.GetWeather)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
}
