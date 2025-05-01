package config

import (
	"github.com/darkphotonKN/flow/internal/booking"
	"github.com/darkphotonKN/flow/internal/user"
	"github.com/gin-gonic/gin"
)

/**
* Sets up API prefix route and all routers.
**/
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// base route
	api := router.Group("/api")

	// -- USER --

	// --- User Setup ---
	userRepo := user.NewUserRepository(DB)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// --- User Routes ---
	userRoutes := api.Group("/user")
	userRoutes.GET("/:id", userHandler.GetUserByIdHandler)
	userRoutes.GET("/", userHandler.GetAllUsersHandler)
	userRoutes.POST("/signup", userHandler.CreateUserHandler)
	userRoutes.POST("/signin", userHandler.LoginUserHandler)

	// -- BOOKING --

	// --- Booking Setup ---
	bookingRepo := booking.NewBookingRepository(DB)
	bookingService := booking.NewBookingService(bookingRepo)
	bookingHandler := booking.NewBookingHandler(bookingService)

	// ---  Booking Routes ---
	bookingRoutes := api.Group("/booking")
	bookingRoutes.POST("/:user_id", bookingHandler.CreateBookingHandler)
	bookingRoutes.GET("/:id", bookingHandler.GetBookingByIdHandler)

	return router
}
