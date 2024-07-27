package models

import (
	"gorm.io/gorm"
)

type Astronaut struct {
	gorm.Model
	Name     string `gorm:"not null;type:varchar(100)"`
	Age      int    `gorm:"not null;type:int"`
	PlanetID uint   `gorm:"not null;type:int"`
	Missions int    `gorm:"not null;type:int"`
	Hours    int    `gorm:"not null;type:int"`
}
