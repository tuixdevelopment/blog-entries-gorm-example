package planet

import (
	"stellar_backend/internal/models"
	"stellar_backend/internal/planet/repository"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(server *echo.Echo, db *gorm.DB) {
	controller := NewPlanetController(db)
	server.GET("/planets", controller.getPlanets)
	server.POST("/planets", controller.createPlanet)
	server.GET("/planets/:id", controller.getPlanet)
	server.PUT("/planets/:id", controller.updatePlanet)
	server.DELETE("/planets/:id", controller.deletePlanet)
}

type planetController struct {
	repository repository.PlanetRepository
}

type PlanetController interface {
	getPlanets(c echo.Context) error
	createPlanet(c echo.Context) error
	getPlanet(c echo.Context) error
	updatePlanet(c echo.Context) error
	deletePlanet(c echo.Context) error
}

func NewPlanetController(db *gorm.DB) PlanetController {
	return &planetController{
		repository: repository.NewPlanetRepository(db),
	}
}

func (controller *planetController) getPlanets(c echo.Context) error {
	planets, err := controller.repository.GetPlanets()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, planets)
}

func (controller *planetController) createPlanet(c echo.Context) error {
	var planet models.Planet
	if err := c.Bind(&planet); err != nil {
		return c.JSON(400, err)
	}
	if err := controller.repository.SavePlanet(&planet); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, planet)

}

func (controller *planetController) getPlanet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	planet, err := controller.repository.GetPlanet(uint(id))
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, planet)
}

func (controller *planetController) updatePlanet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	var planet models.Planet
	if err := c.Bind(&planet); err != nil {
		return c.JSON(400, err)
	}
	planet.ID = uint(id)
	if err := controller.repository.SavePlanet(&planet); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, planet)
}

func (controller *planetController) deletePlanet(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	err = controller.repository.DeletePlanet(uint(id))
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "Planet deleted")
}
