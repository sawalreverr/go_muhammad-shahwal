package controllers

import (
	"eksplorasi_orm_mvc/configs"
	"eksplorasi_orm_mvc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// POST /api/v1/comments/:post_id
func CreateNewCommentController(c echo.Context) error {
	var comment models.Comment
	var resErr models.ResponseError

	id, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	c.Bind(&comment)
	comment.PostID = id

	if createComment := configs.DB.Create(&comment); createComment.Error != nil {
		resErr.Message = createComment.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "COMMENT CREATED",
		Data:    comment,
	}

	return c.JSON(http.StatusOK, resOk)
}

// PUT /api/v1/comments/:id
func EditCommentByIdController(c echo.Context) error {
	var comment models.Comment
	var resErr models.ResponseError

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if commentById := configs.DB.First(&comment, id); commentById.Error != nil {
		resErr.Message = commentById.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	c.Bind(&comment)

	if updateComment := configs.DB.Save(&comment); updateComment.Error != nil {
		resErr.Message = updateComment.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "COMMENT UPDATED",
		Data:    comment,
	}

	return c.JSON(http.StatusOK, resOk)
}

// DELETE /api/v1/comments/:id
func DeleteCommentByIdController(c echo.Context) error {
	var comment models.Comment
	var resErr models.ResponseError

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if commentById := configs.DB.First(&comment, id); commentById.Error != nil {
		resErr.Message = commentById.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	if deleteComment := configs.DB.Delete(&comment, id); deleteComment.Error != nil {
		resErr.Message = deleteComment.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "COMMENT DELETED",
		Data:    "",
	}

	return c.JSON(http.StatusOK, resOk)
}
