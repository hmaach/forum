package models

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	UserName  string
	Content   string
	Likes     int
	Dislikes  int
	CreatedAt string
}
