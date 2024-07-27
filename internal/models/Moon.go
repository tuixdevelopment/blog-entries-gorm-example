package models

import "gorm.io/gorm"

type Moon struct {
	gorm.Model
	Name           string  `gorm:"not null;type:varchar(100)"`
	Description    string  `gorm:"not null;type:varchar(255)"`
	PlanetID       uint    `gorm:"not null;type:int"`
	Mass           float64 `gorm:"not null;type:float"`
	Diameter       float64 `gorm:"not null;type:float"`
	Gravity        float64 `gorm:"not null;type:float"`
	OrbitPeriod    float64 `gorm:"not null;type:float"`
	RotationPeriod float64 `gorm:"not null;type:float"`
	HasAtmosphere  bool    `gorm:"not null;type:boolean"`
}
