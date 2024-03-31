package routes

import (
	"eksplorasi_orm_mvc/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "ON",
		})
	})

	// POST
	e.GET("/api/v1/posts", controllers.GetAllPostController)
	e.GET("/api/v1/posts/:id", controllers.GetPostByIdController)
	e.POST("/api/v1/posts", controllers.CreateNewPostController)
	e.PUT("/api/v1/posts/:id", controllers.EditPostByIdController)
	e.DELETE("/api/v1/posts/:id", controllers.DeletePostByIdController)

	// COMMENT
	e.POST("/api/v1/comments/:post_id", controllers.CreateNewCommentController)
	e.PUT("/api/v1/comments/:id", controllers.EditCommentByIdController)
	e.DELETE("/api/v1/comments/:id", controllers.DeleteCommentByIdController)
}
