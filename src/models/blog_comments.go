package models

type Blog_comments struct {
	Blog_comment_id int
	Blog_id         int
	User_id         int
	Parent_id       int

	Comment      string
	Likes        string
	Added_date   string
	Updated_date string
}
