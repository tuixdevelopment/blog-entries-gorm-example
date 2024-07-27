package repository

import (
	"stellar_backend/internal/models"

	"gorm.io/gorm"
)

type AsteroidRepository interface {
	GetAsteroids() ([]models.Asteroid, error)
	SaveAsteroid(*models.Asteroid) error
	GetAsteroid(uint) (models.Asteroid, error)
	DeleteAsteroid(uint) error
}

type asteroidRepository struct {
	db *gorm.DB
}

func NewAsteroidRepository(db *gorm.DB) AsteroidRepository {
	return &asteroidRepository{db: db}
}

func (r *asteroidRepository) GetAsteroids() ([]models.Asteroid, error) {
	var asteroids []models.Asteroid
	if err := r.db.Find(&asteroids).Error; err != nil {
		return nil, err
	}
	return asteroids, nil
}

func (r *asteroidRepository) SaveAsteroid(asteroid *models.Asteroid) error {
	if err := r.db.Create(asteroid).Error; err != nil {
		return err
	}
	return nil
}

func (r *asteroidRepository) GetAsteroid(id uint) (models.Asteroid, error) {
	var asteroid models.Asteroid
	if err := r.db.First(&asteroid, id).Error; err != nil {
		return models.Asteroid{}, err
	}
	return asteroid, nil
}

func (r *asteroidRepository) DeleteAsteroid(id uint) error {
	if err := r.db.Delete(&models.Asteroid{}, id).Error; err != nil {
		return err
	}
	return nil
}
