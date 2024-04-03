package controllers

import (
	"eksplorasi_middleware/configs"
	"eksplorasi_middleware/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GET /api/v1/posts
func GetAllPostController(c echo.Context) error {
	var posts []models.Post
	var resErr models.ResponseError

	if allPost := configs.DB.Find(&posts); allPost.Error != nil {
		resErr.Message = allPost.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    posts,
	}

	return c.JSON(http.StatusOK, resOk)
}

// GET /api/v1/posts/:id
func GetPostByIdController(c echo.Context) error {
	var post models.Post
	var resErr models.ResponseError

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if postById := configs.DB.First(&post, id); postById.Error != nil {
		resErr.Message = postById.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    post,
	}

	return c.JSON(http.StatusOK, resOk)
}

// POST /api/v1/posts
func CreateNewPostController(c echo.Context) error {
	var post models.Post
	var resErr models.ResponseError

	c.Bind(&post)
	if createPost := configs.DB.Create(&post); createPost.Error != nil {
		resErr.Message = createPost.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "POST CREATED",
		Data:    post,
	}

	return c.JSON(http.StatusOK, resOk)
}

// PUT /api/v1/posts/:id
func EditPostByIdController(c echo.Context) error {
	var post models.Post
	var resErr models.ResponseError

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if postById := configs.DB.First(&post, id); postById.Error != nil {
		resErr.Message = postById.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	c.Bind(&post)

	if updatePost := configs.DB.Save(&post); updatePost.Error != nil {
		resErr.Message = updatePost.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "POST UPDATED",
		Data:    post,
	}

	return c.JSON(http.StatusOK, resOk)
}

// DELETE /api/v1/posts/:id
func DeletePostByIdController(c echo.Context) error {
	var post models.Post
	var resErr models.ResponseError

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if postById := configs.DB.First(&post, id); postById.Error != nil {
		resErr.Message = postById.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	if deletePost := configs.DB.Delete(&post, id); deletePost.Error != nil {
		resErr.Message = deletePost.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "POST DELETED",
		Data:    "",
	}

	return c.JSON(http.StatusOK, resOk)
}
