package wishlist

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Wishlist, error)
	Create(wishlist Wishlist) (Wishlist, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Wishlist, error) {
	var wishlists []Wishlist

	err := r.db.Find(&wishlists).Error

	return wishlists, err
}

func (r *repository) Create(wishlist Wishlist) (Wishlist, error) {
	err := r.db.Create(&wishlist).Error

	return wishlist, err
}
