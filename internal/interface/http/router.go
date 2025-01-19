package http

import (
	"cowork_system/internal/application/usecase/reservation"
	"cowork_system/internal/application/usecase/space"
	"cowork_system/internal/application/usecase/user"
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
	userRepo := repository.NewGormUserRepository(db)
	reservationRepo := repository.NewGormReservationRepository(db)

	//create use cases for spaces
	listSpacesUseCase := space.NewListSpacesUseCase(spaceRepo)
	createSpaceUseCase := space.NewCreateSpaceUseCase(spaceRepo)
	getSpaceByIDUseCase := space.NewGetSpaceUseCase(spaceRepo)
	updateSpaceUseCase := space.NewUpdateSpaceUseCase(spaceRepo)
	deleteSpaceUseCase := space.NewDeleteSpaceUseCase(spaceRepo)

	//create use cases for user
	getUsersUseCase := user.NewGetUsersUseCase(userRepo)
	getUserUseCase := user.NewGetUserUseCase(userRepo)
	createUserUseCase := user.NewCreateUserUseCase(userRepo)
	updateUserUseCase := user.NewUpdateUserUseCase(userRepo)
	deleteUserUseCase := user.NewDeleteUserUseCase(userRepo)
	changeRoleUseCase := user.NewChangeRoleUseCase(userRepo)

	//create use cases for reservation
	getReservationsUseCase := reservation.NewGetReservationsUseCase(reservationRepo)
	getReservationUseCase := reservation.NewGetReservationUseCase(reservationRepo)
	getSpaceReservationsUseCase := reservation.NewGetSpaceReservationsUseCase(reservationRepo)
	getUserReservationsUseCase := reservation.NewGetUserReservationsUseCase(reservationRepo)
	createReservationUseCase := reservation.NewCreateReservationUseCase(reservationRepo, spaceRepo, userRepo)
	updateReservationUseCase := reservation.NewUpdateReservationUseCase(userRepo, spaceRepo, reservationRepo)
	deleteReservationUseCase := reservation.NewDeleteReservationUseCase(reservationRepo)

	//create handler for spaces
	spaceHandler := handler.NewSpaceHandler(
		createSpaceUseCase,
		listSpacesUseCase,
		getSpaceByIDUseCase,
		updateSpaceUseCase,
		deleteSpaceUseCase,
	)

	//create handler for users
	userHandler := handler.NewUserHandler(
		createUserUseCase,
		getUsersUseCase,
		getUserUseCase,
		updateUserUseCase,
		deleteUserUseCase,
		changeRoleUseCase,
	)

	//create handler for reservations
	reservationHandler := handler.NewReservationHandler(
		createReservationUseCase,
		getReservationsUseCase,
		getReservationUseCase,
		getSpaceReservationsUseCase,
		getUserReservationsUseCase,
		updateReservationUseCase,
		deleteReservationUseCase,
	)

	//create routes for spaces
	spaceRoutes := r.Group("/spaces")
	{
		spaceRoutes.GET("/", spaceHandler.GetSpaces)          
		spaceRoutes.POST("/", spaceHandler.CreateSpace)      
		spaceRoutes.GET("/:id", spaceHandler.GetSpaceByID)   
		spaceRoutes.PUT("/:id", spaceHandler.UpdateSpace)   
		spaceRoutes.DELETE("/:id", spaceHandler.DeleteSpace) 
	}

	//create routes for users
	usersRoutes := r.Group("/users")
	{
		usersRoutes.GET("/", userHandler.GetUsers)
		usersRoutes.POST("/", userHandler.CreateUser)
		usersRoutes.GET("/:id", userHandler.GetUser)
		usersRoutes.PUT("/:id", userHandler.UpdateUser)
		usersRoutes.DELETE("/:id", userHandler.DeleteUser)
		usersRoutes.PUT("/role", userHandler.ChangeRole)
	}

	reservationRoutes := r.Group("/reservations")
	{
		reservationRoutes.GET("/", reservationHandler.GetReservations)
		reservationRoutes.POST("/", reservationHandler.CreateReservation)
		reservationRoutes.GET("/:id", reservationHandler.GetReservation)
		reservationRoutes.GET("/space/:id", reservationHandler.GetSpaceReservations)
		reservationRoutes.GET("/user/:id", reservationHandler.GetUserReservations)
		reservationRoutes.PUT("/:id", reservationHandler.UpdateReservation)
		reservationRoutes.DELETE("/:id", reservationHandler.DeleteReservation)
	}

	return r
}