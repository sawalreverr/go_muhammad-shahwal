package models

type Post struct {
	ID       int       `json:"id" form:"id"`
	Title    string    `json:"title" form:"title"`
	Content  string    `json:"content" form:"content"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
