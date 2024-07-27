package models

import (
	"gorm.io/gorm"
)

type Planet struct {
	gorm.Model
	Name              string      `gorm:"not null;type:varchar(100)"`
	Description       string      `gorm:"not null;type:varchar(255)"`
	Moons             []Moon      `gorm:"foreignKey:PlanetID"`
	Mass              float64     `gorm:"not null;type:float"`
	Diameter          float64     `gorm:"not null;type:float"`
	Gravity           float64     `gorm:"not null;type:float"`
	OrbitPeriod       float64     `gorm:"not null;type:float"`
	RotationPeriod    float64     `gorm:"not null;type:float"`
	HasRings          bool        `gorm:"not null;type:boolean"`
	Astronauts        []Astronaut `gorm:"foreignKey:PlanetID"`
	PlanetarySystemID uint        `gorm:"not null;type:int"`
}
