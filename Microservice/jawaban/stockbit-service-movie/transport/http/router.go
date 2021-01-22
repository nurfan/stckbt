package http

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(conn *sqlx.DB) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	adapter := NewAdapter(conn)

	// Routes
	e.GET("/api/v1/movie", adapter.SearchMovie)
	e.GET("/api/v1/movie/:imdbid", adapter.DetailMovie)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
