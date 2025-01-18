package main

import (
	"cowork_system/internal/application/usecase/space"
	"cowork_system/internal/domain/entity"
	"cowork_system/internal/infrastructure/repository"
	"cowork_system/internal/interface/handler"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require TimeZone=Asia/Shanghai", host, user, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	err = db.AutoMigrate(&entity.Space{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	spaceRepo := repository.NewGormSpaceRepository(db)
    listSpacesUseCase := space.NewListSpacesUseCase(spaceRepo)
    createSpaceUseCase := space.NewCreateSpaceUseCase(spaceRepo)

	// Crear handler
	spaceHandler := handler.NewSpaceHandler(createSpaceUseCase, listSpacesUseCase)

	// Crear servidor Gin
	r := gin.Default()

	// Rutas
	r.GET("/spaces", spaceHandler.GetSpaces)
	r.POST("/spaces", spaceHandler.CreateSpace)

	// Iniciar servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}