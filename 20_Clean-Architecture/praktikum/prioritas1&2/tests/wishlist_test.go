package tests

import (
	"encoding/json"
	"go-wishlist-api/controllers"
	"go-wishlist-api/wishlist"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllWishlist(t *testing.T) {
	mockService := &MockWishlistService{}
	expectedWishlists := []wishlist.Wishlist{
		{ID: 1, Title: "Test Wishlist 1", IsAchieved: false},
		{ID: 2, Title: "Test Wishlist 2", IsAchieved: false},
		{ID: 3, Title: "Test Wishlist 3", IsAchieved: true},
	}
	mockService.On("GetAllWishlist").Return(expectedWishlists, nil)

	controller := controllers.NewWishlistController(mockService)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/wishlists", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controller.GetAllWishlist(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	actualWishlist := responseBody["data"].([]interface{})
	convertedWishtlist := make([]wishlist.Wishlist, len(actualWishlist))

	for i, w := range actualWishlist {
		convertedWishtlist[i] = wishlist.Wishlist{
			ID:         uint(w.(map[string]interface{})["id"].(float64)),
			Title:      w.(map[string]interface{})["title"].(string),
			IsAchieved: w.(map[string]interface{})["is_achieved"].(bool),
		}
	}

	assert.Equal(t, expectedWishlists, convertedWishtlist)
}

func TestCreateWishlist(t *testing.T) {
	mockService := &MockWishlistService{}

	expectedWishlists := wishlist.Wishlist{
		ID:         1,
		Title:      "Test Wishlist",
		IsAchieved: false,
	}

	mockService.On("CreateWishlist", mock.AnythingOfType("wishlist.WishlistRequest")).Return(expectedWishlists, nil)

	controller := controllers.NewWishlistController(mockService)

	reqBody := strings.NewReader(`{"title": "Test Wishlist", "is_achieved": false}`)
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/wishlists", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := controller.CreateWishlist(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var responseBody map[string]wishlist.Wishlist
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, expectedWishlists, responseBody["data"])
}
