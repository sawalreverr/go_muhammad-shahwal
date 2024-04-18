package controllers

import (
	"go-wishlist-api/wishlist"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WishtlistController struct {
	wishtlistService wishlist.WishlistService
}

func NewWishlistController(wishlistService wishlist.WishlistService) *WishtlistController {
	return &WishtlistController{wishlistService}
}

func (wc *WishtlistController) GetAllWishlist(c echo.Context) error {
	var wishlists []wishlist.Wishlist

	wishlists, err := wc.wishtlistService.GetAllWishlist()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "data not found",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": wishlists,
	})
}

func (wc *WishtlistController) CreateWishlist(c echo.Context) error {
	var input wishlist.WishlistRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "request invalid",
		})
	}

	wishlist, err := wc.wishtlistService.CreateWishlist(input)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "cannot create data",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"data": wishlist,
	})
}
