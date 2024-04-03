package routes

import (
	"eksplorasi_middleware/controllers"
	"eksplorasi_middleware/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")

	// auth
	v1.POST("/register", controllers.UserRegisterController)
	v1.POST("/login", controllers.UserLoginController)

	// posts (data postingan dapat dikelola oleh user dan admin)
	posts := v1.Group("/posts", middlewares.NewJWTMiddleware(), middlewares.UserValidate)
	posts.GET("", controllers.GetAllPostController)
	posts.GET("/:id", controllers.GetPostByIdController)
	posts.POST("", controllers.CreateNewPostController)
	posts.PUT("/:id", controllers.EditPostByIdController)
	posts.DELETE("/:id", controllers.DeletePostByIdController)

	// categories (data kategori hanya dapat dikelola oleh admin saja)
	categories := v1.Group("/categories", middlewares.NewJWTMiddleware(), middlewares.AdminValidate)
	categories.GET("", controllers.GetAllCategoryController)
	categories.POST("", controllers.CreateNewCategoryController)
	categories.DELETE("", controllers.DeleteCategoryByIdController)
}
