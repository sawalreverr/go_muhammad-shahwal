package wishlist

type WishlistService interface {
	GetAllWishlist() ([]Wishlist, error)
	CreateWishlist(wishlistRequest WishlistRequest) (Wishlist, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllWishlist() ([]Wishlist, error) {
	return s.repository.FindAll()
}

func (s *service) CreateWishlist(wishlistRequest WishlistRequest) (Wishlist, error) {
	wishlist := Wishlist{
		Title:      wishlistRequest.Title,
		IsAchieved: wishlistRequest.IsAchieved,
	}

	return s.repository.Create(wishlist)
}
