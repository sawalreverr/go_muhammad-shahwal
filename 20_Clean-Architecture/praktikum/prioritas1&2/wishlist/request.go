package wishlist

type WishlistRequest struct {
	Title      string `json:"title"`
	IsAchieved bool   `json:"is_achieved"`
}
