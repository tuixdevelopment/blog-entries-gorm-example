package asteroid

import (
	"stellar_backend/internal/asteroid/repository"
	"stellar_backend/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(server *echo.Echo, db *gorm.DB) {
	controller := NewAsteroidController(db)
	server.GET("/asteroids", controller.getAsteroids)
	server.POST("/asteroids", controller.createAsteroid)
	server.GET("/asteroids/:id", controller.getAsteroid)
	server.PUT("/asteroids/:id", controller.updateAsteroid)
	server.DELETE("/asteroids/:id", controller.deleteAsteroid)
}

type asteroidController struct {
	repository repository.AsteroidRepository
}

type AsteroidController interface {
	getAsteroids(c echo.Context) error
	createAsteroid(c echo.Context) error
	getAsteroid(c echo.Context) error
	updateAsteroid(c echo.Context) error
	deleteAsteroid(c echo.Context) error
}

func NewAsteroidController(db *gorm.DB) AsteroidController {
	return &asteroidController{
		repository: repository.NewAsteroidRepository(db),
	}
}

func (controller *asteroidController) getAsteroids(c echo.Context) error {
	asteroids, err := controller.repository.GetAsteroids()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, asteroids)
}

func (controller *asteroidController) createAsteroid(c echo.Context) error {
	var asteroid models.Asteroid
	if err := c.Bind(&asteroid); err != nil {
		return c.JSON(400, err)
	}
	if err := controller.repository.SaveAsteroid(&asteroid); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, asteroid)

}

func (controller *asteroidController) getAsteroid(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	asteroid, err := controller.repository.GetAsteroid(uint(id))
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, asteroid)
}

func (controller *asteroidController) updateAsteroid(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	var asteroid models.Asteroid
	if err := c.Bind(&asteroid); err != nil {
		return c.JSON(400, err)
	}
	asteroid.ID = uint(id)
	if err := controller.repository.SaveAsteroid(&asteroid); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, asteroid)
}

func (controller *asteroidController) deleteAsteroid(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	if err := controller.repository.DeleteAsteroid(uint(id)); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "Asteroid deleted")
}
