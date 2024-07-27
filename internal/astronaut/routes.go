package astronaut

import (
	"stellar_backend/internal/astronaut/repository"
	"stellar_backend/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(server *echo.Echo, db *gorm.DB) {
	controller := NewAstronautController(db)
	server.GET("/astronauts", controller.getAstronauts)
	server.POST("/astronauts", controller.createAstronaut)
	server.GET("/astronauts/:id", controller.getAstronaut)
	server.PUT("/astronauts/:id", controller.updateAstronaut)
	server.DELETE("/astronauts/:id", controller.deleteAstronaut)
}

type astronautController struct {
	repository repository.AstronautRepository
}

type AstronautController interface {
	getAstronauts(c echo.Context) error
	createAstronaut(c echo.Context) error
	getAstronaut(c echo.Context) error
	updateAstronaut(c echo.Context) error
	deleteAstronaut(c echo.Context) error
}

func NewAstronautController(db *gorm.DB) AstronautController {
	return &astronautController{
		repository: repository.NewAstronautRepository(db),
	}
}

func (controller *astronautController) getAstronauts(c echo.Context) error {
	astronauts, err := controller.repository.GetAstronauts()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, astronauts)
}

func (controller *astronautController) createAstronaut(c echo.Context) error {
	var astronaut models.Astronaut
	if err := c.Bind(&astronaut); err != nil {
		return c.JSON(400, err)
	}
	if err := controller.repository.SaveAstronaut(&astronaut); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, astronaut)

}

func (controller *astronautController) getAstronaut(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	astronaut, err := controller.repository.GetAstronaut(uint(id))
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, astronaut)
}

func (controller *astronautController) updateAstronaut(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	var astronaut models.Astronaut
	if err := c.Bind(&astronaut); err != nil {
		return c.JSON(400, err)
	}
	astronaut.ID = uint(id)
	if err := controller.repository.SaveAstronaut(&astronaut); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, astronaut)
}

func (controller *astronautController) deleteAstronaut(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	if err := controller.repository.DeleteAstronaut(uint(id)); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "Astronaut deleted")
}
