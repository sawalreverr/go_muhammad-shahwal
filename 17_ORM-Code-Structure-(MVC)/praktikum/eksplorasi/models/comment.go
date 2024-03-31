package models

type Comment struct {
	ID      int    `json:"id" form:"id"`
	PostID  int    `json:"-" form:"-"`
	Content string `json:"content" form:"content"`
}
