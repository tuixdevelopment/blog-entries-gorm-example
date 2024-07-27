package repository

import (
	"stellar_backend/internal/models"

	"gorm.io/gorm"
)

type PlanetRepository interface {
	GetPlanets() ([]models.Planet, error)
	SavePlanet(*models.Planet) error
	GetPlanet(uint) (models.Planet, error)
	DeletePlanet(uint) error
}

type planetRepository struct {
	db *gorm.DB
}

func NewPlanetRepository(db *gorm.DB) PlanetRepository {
	return &planetRepository{db: db}
}

func (r *planetRepository) GetPlanets() ([]models.Planet, error) {
	var planets []models.Planet
	if err := r.db.Find(&planets).Error; err != nil {
		return nil, err
	}
	return planets, nil
}

func (r *planetRepository) SavePlanet(planet *models.Planet) error {
	if err := r.db.Create(planet).Error; err != nil {
		return err
	}
	return nil
}

func (r *planetRepository) GetPlanet(id uint) (models.Planet, error) {
	var planet models.Planet
	if err := r.db.First(&planet, id).Error; err != nil {
		return models.Planet{}, err
	}
	return planet, nil
}

func (r *planetRepository) DeletePlanet(id uint) error {
	if err := r.db.Delete(&models.Planet{}, id).Error; err != nil {
		return err
	}
	return nil
}
