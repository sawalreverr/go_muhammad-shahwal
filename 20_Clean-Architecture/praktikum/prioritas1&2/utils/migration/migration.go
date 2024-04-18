package migration

import (
	"fmt"
	"go-wishlist-api/user"
	"go-wishlist-api/utils/database"
	"go-wishlist-api/wishlist"
	"log"
)

func AutoMigrate() {
	if err := database.DB.AutoMigrate(&wishlist.Wishlist{}, &user.User{}); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	fmt.Println("Database Migration Success")
}
