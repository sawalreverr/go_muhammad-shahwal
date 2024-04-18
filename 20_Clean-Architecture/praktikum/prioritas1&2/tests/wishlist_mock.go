package tests

import (
	"go-wishlist-api/wishlist"

	"github.com/stretchr/testify/mock"
)

//go:generate mock -source=../wishlist/wishlist.go -destination=wishlist_mock.go -package=tests

type MockWishlistService struct {
	mock.Mock
}

func (m *MockWishlistService) GetAllWishlist() ([]wishlist.Wishlist, error) {
	args := m.Called()
	if res := args.Get(0); res != nil {
		return res.([]wishlist.Wishlist), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockWishlistService) CreateWishlist(wishlist.WishlistRequest) (wishlist.Wishlist, error) {
	args := m.Called(wishlist.WishlistRequest{})
	if res := args.Get(0); res != nil {
		return res.(wishlist.Wishlist), args.Error(1)
	}
	return wishlist.Wishlist{}, args.Error(1)
}
