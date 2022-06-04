package models

type Blogs struct {
	Blog_id          int
	Blog_category_id int
	User_id          int
	Is_popular       int

	Title        string
	Description  string
	Thumbnail    string
	Banner       string
	Likes        string
	Added_date   string
	Updated_date string
	Status       string
	Keywords     string
}
