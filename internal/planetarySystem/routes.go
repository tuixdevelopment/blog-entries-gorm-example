package planetarysystem

import (
	"stellar_backend/internal/models"
	"stellar_backend/internal/planetarySystem/repository"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(server *echo.Echo, db *gorm.DB) {
	controller := NewPlanetarysystemController(db)
	server.GET("/planetarysystems", controller.getPlanetarySystems)
	server.POST("/planetarysystems", controller.createPlanetarySystem)
	server.GET("/planetarysystems/:id", controller.getPlanetarySystem)
	server.PUT("/planetarysystems/:id", controller.updatePlanetarySystem)
	server.DELETE("/planetarysystems/:id", controller.deletePlanetarySystem)
}

type planetarysystemController struct {
	repository repository.PlanetarysystemRepository
}

type PlanetarysystemController interface {
	getPlanetarySystems(c echo.Context) error
	createPlanetarySystem(c echo.Context) error
	getPlanetarySystem(c echo.Context) error
	updatePlanetarySystem(c echo.Context) error
	deletePlanetarySystem(c echo.Context) error
}

func NewPlanetarysystemController(db *gorm.DB) PlanetarysystemController {
	return &planetarysystemController{
		repository: repository.NewPlanetarysystemRepository(db),
	}
}

func (controller *planetarysystemController) getPlanetarySystems(c echo.Context) error {
	systems, err := controller.repository.GetPlanetarySystems()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, systems)
}

func (controller *planetarysystemController) createPlanetarySystem(c echo.Context) error {
	var planetarySystem models.PlanetarySystem
	if err := c.Bind(&planetarySystem); err != nil {
		return c.JSON(400, err)
	}
	if err := controller.repository.SavePlanetarySystem(&planetarySystem); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, planetarySystem)

}

func (controller *planetarysystemController) getPlanetarySystem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	system, err := controller.repository.GetPlanetarySystem(uint(id))
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, system)
}

func (controller *planetarysystemController) updatePlanetarySystem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	var planetarySystem models.PlanetarySystem
	if err := c.Bind(&planetarySystem); err != nil {
		return c.JSON(400, err)
	}
	planetarySystem.ID = uint(id)
	if err := controller.repository.SavePlanetarySystem(&planetarySystem); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, planetarySystem)
}

func (controller *planetarysystemController) deletePlanetarySystem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	if err := controller.repository.DeletePlanetarySystem(uint(id)); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "Planetary System deleted")
}
