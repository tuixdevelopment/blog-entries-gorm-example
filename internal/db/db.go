package db

import (
	"fmt"
	"os"
	"stellar_backend/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=" + os.Getenv("SSL_MODE") +
		" TimeZone=" + os.Getenv("TIMEZONE")

	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database! 🚀")

	db.AutoMigrate(
		&models.PlanetarySystem{},
		&models.Planet{},
		&models.Asteroid{},
		&models.Moon{},
		&models.Astronaut{},
	)
}

func DB() *gorm.DB {
	return db
}
