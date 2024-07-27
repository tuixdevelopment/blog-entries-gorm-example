package repository

import (
	"stellar_backend/internal/models"

	"gorm.io/gorm"
)

type PlanetarysystemRepository interface {
	GetPlanetarySystems() ([]models.PlanetarySystem, error)
	SavePlanetarySystem(*models.PlanetarySystem) error
	GetPlanetarySystem(uint) (models.PlanetarySystem, error)
	DeletePlanetarySystem(uint) error
}

type planetarysystemRepository struct {
	db *gorm.DB
}

func NewPlanetarysystemRepository(db *gorm.DB) PlanetarysystemRepository {
	return &planetarysystemRepository{db: db}
}

func (r *planetarysystemRepository) GetPlanetarySystems() ([]models.PlanetarySystem, error) {
	var planetarySystems []models.PlanetarySystem
	if err := r.db.Find(&planetarySystems).Error; err != nil {
		return nil, err
	}
	return planetarySystems, nil
}

func (r *planetarysystemRepository) SavePlanetarySystem(planetarySystem *models.PlanetarySystem) error {
	if err := r.db.Create(planetarySystem).Error; err != nil {
		return err
	}
	return nil
}

func (r *planetarysystemRepository) GetPlanetarySystem(id uint) (models.PlanetarySystem, error) {
	var planetarySystem models.PlanetarySystem
	if err := r.db.First(&planetarySystem, id).Error; err != nil {
		return models.PlanetarySystem{}, err
	}
	return planetarySystem, nil
}

func (r *planetarysystemRepository) DeletePlanetarySystem(id uint) error {
	if err := r.db.Delete(&models.PlanetarySystem{}, id).Error; err != nil {
		return err
	}
	return nil
}
