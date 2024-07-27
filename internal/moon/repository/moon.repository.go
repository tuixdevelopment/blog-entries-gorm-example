package repository

import (
	"stellar_backend/internal/models"

	"gorm.io/gorm"
)

type MoonRepository interface {
	GetMoons() ([]models.Moon, error)
	SaveMoon(*models.Moon) error
	GetMoon(uint) (models.Moon, error)
	DeleteMoon(uint) error
}

type moonRepository struct {
	db *gorm.DB
}

func NewMoonRepository(db *gorm.DB) MoonRepository {
	return &moonRepository{db: db}
}

func (r *moonRepository) GetMoons() ([]models.Moon, error) {
	var moons []models.Moon
	if err := r.db.Find(&moons).Error; err != nil {
		return nil, err
	}
	return moons, nil
}

func (r *moonRepository) SaveMoon(moon *models.Moon) error {
	if err := r.db.Create(moon).Error; err != nil {
		return err
	}
	return nil
}

func (r *moonRepository) GetMoon(id uint) (models.Moon, error) {
	var moon models.Moon
	if err := r.db.First(&moon, id).Error; err != nil {
		return models.Moon{}, err
	}
	return moon, nil
}

func (r *moonRepository) DeleteMoon(id uint) error {
	if err := r.db.Delete(&models.Moon{}, id).Error; err != nil {
		return err
	}
	return nil
}
