package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osmodiar16/evalua-tu-profe/internal/app/handlers"
)

func main() {
	// Echo instance
	e := echo.New()

	userHandler := handlers.UserHandler{}
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/user", userHandler.HandleUserShow)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Bitchs!")
}
