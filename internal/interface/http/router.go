package http

import (
	"cowork_system/internal/application/usecase/space"
	"cowork_system/internal/infrastructure/database"
	"cowork_system/internal/infrastructure/repository"
	"cowork_system/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db, err := database.NewDBConnection()
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	//create repositories of each entity
	spaceRepo := repository.NewGormSpaceRepository(db)

	//create use cases of each entity
	listSpacesUseCase := space.NewListSpacesUseCase(spaceRepo)
	createSpaceUseCase := space.NewCreateSpaceUseCase(spaceRepo)
	getSpaceByIDUseCase := space.NewGetSpaceUseCase(spaceRepo)
	updateSpaceUseCase := space.NewUpdateSpaceUseCase(spaceRepo)
	deleteSpaceUseCase := space.NewDeleteSpaceUseCase(spaceRepo)

	//create handler of each entity
	spaceHandler := handler.NewSpaceHandler(
		createSpaceUseCase,
		listSpacesUseCase,
		getSpaceByIDUseCase,
		updateSpaceUseCase,
		deleteSpaceUseCase,
	)

	//create routes of each entity
	spaceRoutes := r.Group("/spaces")
	{
		spaceRoutes.GET("/", spaceHandler.GetSpaces)          
		spaceRoutes.POST("/", spaceHandler.CreateSpace)      
		spaceRoutes.GET("/:id", spaceHandler.GetSpaceByID)   
		spaceRoutes.PUT("/:id", spaceHandler.UpdateSpace)   
		spaceRoutes.DELETE("/:id", spaceHandler.DeleteSpace) 
	}

	return r
}