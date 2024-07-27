package models

import "gorm.io/gorm"

type Asteroid struct {
	gorm.Model
	Name              string  `gorm:"not null;type:varchar(100)"`
	Description       string  `gorm:"not null;type:varchar(255)"`
	PlanetarySystemID uint    `gorm:"not null;type:int"`
	Mass              float64 `gorm:"not null;type:float"`
	Diameter          float64 `gorm:"not null;type:float"`
}
