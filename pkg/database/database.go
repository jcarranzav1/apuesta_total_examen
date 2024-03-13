package database

import (
	"fmt"
	"log"

	"ApuestaTotal/config"
	"ApuestaTotal/internal/products/infrastructure/adapters/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	env := config.Environments()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		env.DbHost,
		env.DbUser,
		env.DbPassword,
		env.DbName,
		env.DbPort,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info), // Cambiar a logger.Silent en PROD
	})
	if err != nil {
		log.Fatal("Failed to connect to database: /n ", err)
	}

	log.Println("Database Connected")
	log.Println("Running migrations...")

	_ = database.AutoMigrate(
		&model.Product{},
	)

	return database
}
