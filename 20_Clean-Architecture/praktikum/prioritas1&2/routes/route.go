package routes

import (
	"go-wishlist-api/auth"
	"go-wishlist-api/controllers"
	"go-wishlist-api/middlewares"
	"go-wishlist-api/user"
	"go-wishlist-api/utils/database"
	"go-wishlist-api/wishlist"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	wishlistRepository := wishlist.NewRepository(database.DB)
	wishlistService := wishlist.NewService(wishlistRepository)
	wishlistController := controllers.NewWishlistController(wishlistService)

	userRepository := user.NewRepository(database.DB)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userController := controllers.NewUserController(userService, authService)

	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)

	// need authentication
	auth := e.Group("/auth")
	auth.Use(middlewares.AuthMiddleware(authService, userService))

	auth.GET("/wishlists", wishlistController.GetAllWishlist)
	auth.POST("/wishlists", wishlistController.CreateWishlist)
}
