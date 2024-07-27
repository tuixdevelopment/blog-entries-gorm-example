package server

import (
	"stellar_backend/internal/asteroid"
	"stellar_backend/internal/astronaut"
	"stellar_backend/internal/db"
	"stellar_backend/internal/moon"
	"stellar_backend/internal/planet"
	planetarysystem "stellar_backend/internal/planetarySystem"

	"github.com/labstack/echo/v4"
)

var Server *echo.Echo

func init() {
	Server = echo.New()

	db := db.DB()
	planet.InitRoutes(Server, db)
	asteroid.InitRoutes(Server, db)
	astronaut.InitRoutes(Server, db)
	planetarysystem.InitRoutes(Server, db)
	moon.InitRoutes(Server, db)
}
