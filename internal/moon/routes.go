package moon

import (
	"stellar_backend/internal/models"
	"stellar_backend/internal/moon/repository"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(server *echo.Echo, db *gorm.DB) {
	controller := NewMoonController(db)
	server.GET("/moon", controller.getMoons)
	server.POST("/moon", controller.createMoon)
	server.GET("/moon/:id", controller.getMoon)
	server.PUT("/moon/:id", controller.updateMoon)
	server.DELETE("/moon/:id", controller.deleteMoon)
}

type moonController struct {
	repository repository.MoonRepository
}

type MoonController interface {
	getMoons(c echo.Context) error
	createMoon(c echo.Context) error
	getMoon(c echo.Context) error
	updateMoon(c echo.Context) error
	deleteMoon(c echo.Context) error
}

func NewMoonController(db *gorm.DB) MoonController {
	return &moonController{
		repository: repository.NewMoonRepository(db),
	}
}

func (controller *moonController) getMoons(c echo.Context) error {
	moons, err := controller.repository.GetMoons()
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, moons)
}

func (controller *moonController) createMoon(c echo.Context) error {
	var moon models.Moon
	if err := c.Bind(&moon); err != nil {
		return c.JSON(400, err)
	}
	if err := controller.repository.SaveMoon(&moon); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(201, moon)

}

func (controller *moonController) getMoon(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	moon, err := controller.repository.GetMoon(uint(id))
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, moon)
}

func (controller *moonController) updateMoon(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	var moon models.Moon
	if err := c.Bind(&moon); err != nil {
		return c.JSON(400, err)
	}
	moon.ID = uint(id)
	if err := controller.repository.SaveMoon(&moon); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, moon)
}

func (controller *moonController) deleteMoon(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	if err := controller.repository.DeleteMoon(uint(id)); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "Moon deleted")
}
