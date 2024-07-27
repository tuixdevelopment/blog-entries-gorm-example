package repository

import (
	"stellar_backend/internal/models"

	"gorm.io/gorm"
)

type AstronautRepository interface {
	GetAstronauts() ([]models.Astronaut, error)
	SaveAstronaut(*models.Astronaut) error
	GetAstronaut(uint) (models.Astronaut, error)
	DeleteAstronaut(uint) error
}

type astronautRepository struct {
	db *gorm.DB
}

func NewAstronautRepository(db *gorm.DB) AstronautRepository {
	return &astronautRepository{db: db}
}

func (r *astronautRepository) GetAstronauts() ([]models.Astronaut, error) {
	var astronauts []models.Astronaut
	if err := r.db.Find(&astronauts).Error; err != nil {
		return nil, err
	}
	return astronauts, nil
}

func (r *astronautRepository) SaveAstronaut(astronaut *models.Astronaut) error {
	if err := r.db.Create(astronaut).Error; err != nil {
		return err
	}
	return nil
}

func (r *astronautRepository) GetAstronaut(id uint) (models.Astronaut, error) {
	var astronaut models.Astronaut
	if err := r.db.First(&astronaut, id).Error; err != nil {
		return models.Astronaut{}, err
	}
	return astronaut, nil
}

func (r *astronautRepository) DeleteAstronaut(id uint) error {
	if err := r.db.Delete(&models.Astronaut{}, id).Error; err != nil {
		return err
	}
	return nil
}
