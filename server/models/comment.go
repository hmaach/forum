package models

type Comment struct {
	ID       int
	UserID   int
	PostID   int
	Content  string
	CreatedAt string
}
