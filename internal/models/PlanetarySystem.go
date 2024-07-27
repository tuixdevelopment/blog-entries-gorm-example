package models

import "gorm.io/gorm"

type PlanetarySystem struct {
	gorm.Model
	Name      string     `gorm:"not null;type:varchar(100)" json:"name"`
	Planets   []Planet   `gorm:"foreignKey:PlanetarySystemID" json:"planets"`
	Asteroids []Asteroid `gorm:"foreignKey:PlanetarySystemID" json:"asteroids"`
}
